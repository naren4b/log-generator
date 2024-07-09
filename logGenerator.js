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

// Get the duration for log generation in milliseconds from environment variable or default to 5 minutes
const logDuration = process.env.LOG_DURATION || (5 * 60 * 1000); // 5 minutes in milliseconds

// Calculate end time for log generation
const endTime = Date.now() + logDuration;

// Generate logs until endTime is reached
function generateLogsUntilEnd() {
  if (Date.now() < endTime) {
    generateRandomLog();
    setTimeout(generateLogsUntilEnd, Math.random() * 2000); // Random delay up to 2 seconds
  } else {
    console.log('Log generation completed.');
    process.exit(0);
  }
}

// Start generating logs
generateLogsUntilEnd();
