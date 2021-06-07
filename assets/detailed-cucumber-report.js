// documentation:
//
const report = require('multiple-cucumber-html-reporter');

var os = require('os');

let report_location = "test-results/"

report.generate({
    jsonDir: report_location,
    reportPath: report_location,
    reportName: 'Cucumber detailed report',
    displayReportTime: false,
    hideMetadata: true,
    openReportInBrowser: true,
    metadata:{
        "Platform": os.platform() + "-" + os.release(),
        "Executed": new Date(),

        platform: {
            name: os.platform() + "-" + os.release(),

        }
    }
});

