const fs = require('fs');
const path = require('path');
const rcedit = require('rcedit');

const exePath = path.join(__dirname, 'bin/double-tap.exe');
const iconPath = path.join(__dirname, 'resources/win32/icon.ico');

(async () => {
  if (fs.existsSync(exePath)) {
    await rcedit(exePath, {
      icon: iconPath
    });
  }
})();
