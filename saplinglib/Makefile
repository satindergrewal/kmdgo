# ifeq ($(shell uname),Darwin)
#     EXT := dylib
# else
#     EXT := a
# endif
EXT := a

all: target/release/saplinglib.$(EXT)

target/release/saplinglib.$(EXT): src/lib.rs Cargo.toml
	cargo build --lib --release
	@cp -av target/release/libsaplinglib.$(EXT) .
	@rm -rf target

cbin: target/release/saplinglib.$(EXT)
	gcc src/main.c -L./ -lsaplinglib -lpthread -ldl -o csapling -framework Security
	./csapling
	@rm -rf target
	@rm -rf Cargo.lock

gobin: target/release/saplinglib.$(EXT)
	go build -o gosapling src/main.go
	./gosapling
	@rm -rf target
	@rm -rf Cargo.lock

clean:
	@rm -rf target
	@rm -rf csapling
	@rm -rf gosapling
	@rm -rf Cargo.lock
	@rm -rf libsaplinglib.$(EXT)