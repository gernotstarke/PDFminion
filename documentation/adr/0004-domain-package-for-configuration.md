# 4. domain-package-for-configuration

Date: 2024-12-01

## Status

Accepted

## Context

PDFminion needs a number of configuration options to be set by the user.
These options are used to control the behavior of the application.
The configuration options are used in multiple places in the application, and it is important that they are consistent.

## Decision

We will create a domain package to hold configuration types and functions.

## Consequences

* We need to handle the dependencies between the cli package and the domain package (the cli package needs to know about the configuration types).