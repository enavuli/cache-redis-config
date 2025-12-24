const fs = require('fs');

module.exports = {
  parse: (filePath) => {
    const fileContent = fs.readFileSync(filePath, 'utf8').split('\n');
    const lines = [];
    let currentSection = false;

    fileContent.forEach(line => {
      const trimmedLine = line.trim();

      if (trimmedLine.startsWith('[') && trimmedLine.endsWith(']')) {
        currentSection = trimmedLine.replace('[', '').replace(']', '');
        lines.push({ section: currentSection, items: [] });
      } else if (trimmedLine && trimmedLine.length > 0 && currentSection) {
        lines[lines.length - 1].items.push(trimmedLine);
      }
    });

    const parsedConfig = lines.reduce((acc, line) => {
      if (line.section) {
        acc[line.section] = line.items;
      }
      return acc;
    }, {});

    return parsedConfig;
  }
};