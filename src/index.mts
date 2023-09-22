import ffi from "ffi-napi";
import ref from "ref-napi";

const callbackFunc = ffi.Callback("int", ["int"], function (n: number) {
  console.log(n);
});

const lib = ffi.Library("./native-artifacts/node-libuiohook-linux-amd64.so", {
  Start: ["void", []],
  End: ["void", []],
  SetReceiver: ["void", ["pointer"]]
});

lib.Start();

setInterval(() => {
  // lib.End();
}, 5_000);