const { getDefaultConfig } = require("@expo/metro-config");
const path = require("path");

const projectRoot = __dirname;

// Create the default Metro config
const config = getDefaultConfig(projectRoot);

// Add import aliases
config.resolver.alias = {
    "@screens": path.resolve(projectRoot, "screens"),
    "@assets": path.resolve(projectRoot, "assets"),
    "@app": path.resolve(projectRoot, "app"),
};

module.exports = config;
