const bcrypt = require("bcrypt");

const password = "1234Segura!";

bcrypt.hash(password, 10, (err, hash) => {
    console.log(hash);

    bcrypt.compare(password, hash, (err, res) =>{
        console.log(res);
    })
})