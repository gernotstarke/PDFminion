var reporter = require('cucumber-html-reporter');

var os = require('os');

// file location
let report_location = "test-results/"

//var reporter_theme = 'hierarchy'
var reporter_theme = 'bootstrap'

var options = {
    theme: reporter_theme,
    jsonFile: report_location + 'cucumber-report.json',
    output: report_location + 'simple-cucumber-report.html',
    brandTitle: 'PDFminion Cucumber Report',
    reportSuiteAsScenarios: true,
    scenarioTimestamp: true,
    launchReport: true,
    metadata: {
        "Platform": os.platform() + "-" + os.release(),
        "Executed": new Date(),
        "Theme": reporter_theme
    }
};

reporter.generate(options);

