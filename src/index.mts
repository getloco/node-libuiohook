import koffi from "koffi";

const lib = koffi.load("./native-artifacts/node-libuiohook-linux-amd64.so");

const listener = koffi.proto(`
  void listener(
    uint8_t kind,
    int64_t when,
    uint16_t mask,
    uint16_t reserved,
    uint16_t key_code,
    uint16_t raw_code,
    int32_t key_char,
    uint16_t button,
    uint16_t clicks,
    int16_t x,
    int16_t y,
    uint16_t amount,
    int32_t rotaion,
    uint8_t direction
  )
`);

const Join = lib.func("Join", "void", [koffi.pointer(listener)]);

Join.async(
  (...args: any[]) => {
    console.log(args);
  },
  () => {}
);
