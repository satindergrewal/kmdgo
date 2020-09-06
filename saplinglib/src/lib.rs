use libc::{c_char};
use std::ffi::{CStr, CString};
use saplingrustlib;

/**
 * Call into rust to generate a paper wallet. Returns the paper wallet in JSON form. 
 * NOTE: the returned string is owned by rust, so the caller needs to call rust_free_string with it
 * after using it to free it properly
 */ 
#[no_mangle]
pub extern fn rust_generate_wallet(nohd: bool, zcount: u32, entropy: *const c_char, iguana_seed: bool) -> *mut c_char {
    let entropy_str = unsafe {
        assert!(!entropy.is_null());

        CStr::from_ptr(entropy)
    };

    let c_str = CString::new(saplingrustlib::generate_wallet(nohd, zcount, entropy_str.to_bytes(), iguana_seed)).unwrap();
    return c_str.into_raw();
}

/**
 * Callers that recieve string return values from other functions should call this to return the string 
 * back to rust, so it can be freed. Failure to call this function will result in a memory leak
 */ 
#[no_mangle]
pub extern fn rust_free_string(s: *mut c_char) {
    unsafe {
        if s.is_null() { return }
        CString::from_raw(s)
    };
}
