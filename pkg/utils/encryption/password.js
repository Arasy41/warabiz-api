const bcrypt = require('bcryptjs');

function EncryptPassword(password) {
    const salt = bcrypt.genSaltSync();
    const hashedPassword = bcrypt.hashSync(password, salt);
    return hashedPassword;
}
