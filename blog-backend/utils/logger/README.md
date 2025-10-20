# Logger Utility

This package provides a structured logging solution using Logrus for the blog backend application.

## Usage

To use the logger in any Go file, import the package:

```go
import "com.tang.blog/utils/logger"
```

Then you can use the logger like this:

```go
// Simple info log
logger.Log.Info("Application started")

// Error log with fields
logger.Log.WithFields(map[string]interface{}{
    "error": err.Error(),
    "userID": userID,
}).Error("Failed to process user request")

// Debug log
logger.Log.WithFields(map[string]interface{}{
    "postID": post.ID,
    "title": post.Title,
}).Debug("Processing post")
```

## Log Levels

The logger supports the following log levels:
- Debug
- Info
- Warning
- Error
- Fatal
- Panic

## Log Output

Logs are written to both the console and to daily log files in the `logs/` directory. Each log file is named with the date (e.g., `2023-10-17.log`).

## Log Format

Logs are formatted with timestamps in RFC3339 format for better readability and parsing.