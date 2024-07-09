const winston = require('winston');

// Custom log levels
const customLevels = {
  levels: {
    info: 0,
    debug: 1,
    audit: 2,
    trace: 3,
    warn: 4,
    error: 5
  },
  colors: {
    info: 'green',
    debug: 'blue',
    audit: 'magenta',
    trace: 'cyan',
    warn: 'yellow',
    error: 'red'
  }
};

// Create a Winston logger with the custom levels
const logger = winston.createLogger({
  levels: customLevels.levels,
  format: winston.format.combine(
    winston.format.colorize(),
    winston.format.timestamp(),
    winston.format.printf(({ timestamp, level, message }) => `${timestamp} [${level}]: ${message}`)
  ),
  transports: [
    new winston.transports.Console()
  ]
});

// Apply colors to the custom levels
winston.addColors(customLevels.colors);

// Simulated users
const users = ['User1', 'User2', 'User3'];

// Simulated actions
const actions = [
  'login',
  'logout',
  'update inventory',
  'delete inventory'
];

// Function to generate a random log message
function generateRandomLog() {
  const randomUser = users[Math.floor(Math.random() * users.length)];
  const randomAction = actions[Math.floor(Math.random() * actions.length)];
  const logMessage = `${randomUser} performed ${randomAction}.`;

  logger.log('audit', logMessage);
}

// Get the log generation interval from the environment variable or default to 1000 ms
const interval = process.env.LOG_INTERVAL || 1000;

// Generate a random log at the specified interval
setInterval(generateRandomLog, interval);

console.log(`Audit log generator started. It will generate a random log every ${interval} milliseconds.`);
