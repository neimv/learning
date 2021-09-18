const os = require("os");

console.log(os.arch());
console.log(os.platform());
console.log(os.cpus().length);
console.log(os.constants);
console.log(os.freemem());

console.log(os.homedir());
console.log(os.tmpdir());
console.log(os.hostname());
console.log(os.networkInterfaces());