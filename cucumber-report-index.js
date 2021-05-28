var reporter = require('cucumber-html-reporter');

var os = require('os');

//var reporter_theme = 'hierarchy'
var reporter_theme = 'bootstrap'

var options = {
    theme: reporter_theme,
    jsonFile: 'test-results/cucumber-report.json',
    output: 'test-results/cucumber_report.html',
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

