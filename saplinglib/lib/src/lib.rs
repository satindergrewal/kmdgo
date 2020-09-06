use hex;
use zip32::{ChildIndex, ExtendedSpendingKey};
use bech32::{Bech32, u5, ToBase32};
use rand::{Rng, ChaChaRng, FromEntropy, SeedableRng};
use json::{array, object};
use blake2_rfc::blake2b::Blake2b;

/**
 * Generate a series of `zcount` addresses and private keys. 
 */
pub fn generate_wallet(nohd: bool, zcount: u32, user_entropy: &[u8], iguana_seed: bool) -> String {        
    // Get 32 bytes of system entropy
    let mut system_entropy:[u8; 32] = [0; 32]; 
    {
        let mut system_rng = ChaChaRng::from_entropy();    
        system_rng.fill(&mut system_entropy);
    }

    // Add in user entropy to the system entropy, and produce a 32 byte hash... 
    let mut state = Blake2b::new(32);
    // if iguana_seed is set to true, ignore using system entropy, so that a deterministic address is generated
    if !iguana_seed {
        // iguana_seed is false. use system entropy as usual
        state.update(&system_entropy);
    }
    state.update(&user_entropy);
    
    let mut final_entropy: [u8; 32] = [0; 32];
    final_entropy.clone_from_slice(&state.finalize().as_bytes()[0..32]);

    // ...which will we use to seed the RNG
    let mut rng = ChaChaRng::from_seed(final_entropy);

    if !nohd {
        // Allow HD addresses, so use only 1 seed
        let mut seed: [u8; 32] = [0; 32];
        rng.fill(&mut seed);
        
        return gen_addresses_with_seed_as_json(zcount, |i| (seed.to_vec(), i));
    } else {
        // Not using HD addresses, so derive a new seed every time    
        return gen_addresses_with_seed_as_json(zcount, |_| {            
            let mut seed:[u8; 32] = [0; 32]; 
            rng.fill(&mut seed);
            
            return (seed.to_vec(), 0);
        });
    }    
}

/**
 * Generate `zcount` addresses with the given seed. The addresses are derived from m/32'/cointype'/index' where 
 * index is 0..zcount
 * 
 * Note that cointype is 1 for testnet and 133 for mainnet
 * 
 * get_seed is a closure that will take the address number being derived, and return a tuple cointaining the 
 * seed and child number to use to derive this wallet. 
 *
 * It is useful if we want to reuse (or not) the seed across multiple wallets.
 */
fn gen_addresses_with_seed_as_json<F>(zcount: u32, mut get_seed: F) -> String 
    where F: FnMut(u32) -> (Vec<u8>, u32)
{
    let mut ans = array![];

    for i in 0..zcount {
        let (seed, child) = get_seed(i);
        let (addr, pk, path) = get_address(&seed, child);
        ans.push(object!{
                "num"           => i,
                "address"       => addr,
                "private_key"   => pk,
                "seed"          => path
        }).unwrap(); 
    }      

    return json::stringify_pretty(ans, 2);
}

// Generate a standard ZIP-32 address from the given seed at 32'/44'/0'/index
fn get_address(seed: &[u8], index: u32) -> (String, String, json::JsonValue) {
    let addr_prefix : &str = "zs";
    let pk_prefix   : &str = "secret-extended-key-main";
    let cointype           = {133};
    
    let spk: ExtendedSpendingKey = ExtendedSpendingKey::from_path(
            &ExtendedSpendingKey::master(seed),
            &[
                ChildIndex::Hardened(32),
                ChildIndex::Hardened(cointype),
                ChildIndex::Hardened(index)
            ],
        );
    let path = object!{
        "HDSeed"    => hex::encode(seed),
        "path"      => format!("m/32'/{}'/{}'", cointype, index)
    };

    let (_d, addr) = spk.default_address().expect("Cannot get result");

    // Address is encoded as a bech32 string
    let mut v = vec![0; 43];
    v.get_mut(..11).unwrap().copy_from_slice(&addr.diversifier.0);
    addr.pk_d.write(v.get_mut(11..).unwrap()).expect("Cannot write!");
    let checked_data: Vec<u5> = v.to_base32();
    let encoded = Bech32::new(addr_prefix.into(), checked_data).expect("bech32 failed").to_string();

    // Private Key is encoded as bech32 string
    let mut vp = Vec::new();
    spk.write(&mut vp).expect("Can't write private key");
    let c_d: Vec<u5> = vp.to_base32();
    let encoded_pk = Bech32::new(pk_prefix.into(), c_d).expect("bech32 failed").to_string();

    return (encoded.to_string(), encoded_pk.to_string(), path);
}



