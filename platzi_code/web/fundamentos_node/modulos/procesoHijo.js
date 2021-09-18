const { exec, spawn } = require("child_process");

exec("ls -la", (err, stdout, stderr) => {
    if (err) {
        console.error(err);
        return false;
    }

    console.log(stdout);
});

let proceso = spawn("ls", ["-la"]);

console.log(proceso.pid);
console.log(proceso.connected);

proceso.stdout.on("data", (dato) => { console.log(dato.toString()); });

proceso.on("exit", () => { console.log("el proceso termino"); });
