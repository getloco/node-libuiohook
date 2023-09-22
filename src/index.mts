import ffi from "ffi-napi";
import ref from "ref-napi";

let i = 0;

const callbackFunc = ffi.Callback(
  "void",
  [
    "uint8", // kind
    "int64", // when
    "uint16", // mask
    "uint16", // reserved
    "uint16", // key_code
    "uint16", // raw_code
    "int32", // key_char
    "uint16", // button
    "uint16", // clicks
    "int16", // x
    "int16", // y
    "uint16", // amount
    "int32", // rotation
    "uint8", // direction
  ],
  function (...args) {
    console.log(i++, args);
  }
);

const lib = ffi.Library("./native-artifacts/node-libuiohook-linux-amd64.so", {
  Join: ["void", ["pointer"]],
  Kill: ["void", []],
});

lib.Join.async(callbackFunc, () => {});

setInterval(() => {
  // console.log("huehue");
}, 5_000);
