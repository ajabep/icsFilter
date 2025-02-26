# ICS Filter

[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/ajabep/icsFilter/badge)](https://securityscorecards.dev/viewer/?uri=github.com/ajabep/icsFilter)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=ajabep_icsFilter&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=ajabep_icsFilter)

ICSFilter is a (quick and dirty) simple HTTP server to filter iCal (.ics) URL.

I create it because I have third parties who are exporting calendars in a iCal (.ics) address. I add them to my 
calendar. But, the flow is dirty, and they don't allow me to filter event during the export.

By running this tool, I create an HTTP server, acting as a proxy to the third party subscriptions and apply the 
filters each time my calendar wants to refresh the flow.

This server is not made for being blazing fast. Just to work and not being a pain to code or deploy.

## How to run?

### Docker (Recommended)

For the configuration, please, refer to the "How to configure" section.

TODO

### By a running command line

This method is not recommended. Indeed, the server will work while the command is running. Thus, if the command is 
run in foreground, you will only get your terminal once the server will be stopped.

Also, this requires from you to have installed Go.

To install, just run:

```bash
go install github.com/ajabep/icsFiltercmd/icsfilter@latest
```

For the configuration, please, refer to the "How to configure" section.

To run the server in foreground (will block your terminal), run:

```bash
icsfilter ./configurationFile.yml
```

The HTTP server will being exposed on the port tcp/8080 on all your interfaces. The routing is up to you. If you're 
using the command line, you should know what you are doing. No HTTPS usage is supported (yet), you need to use a 
reverse proxy for now.

## How to configure?

An example of the configuration is present at [./rule.yml.sample](./rule.yml.sample).

The configuration is though as described in the following subsections.

To validate the configuration, run the server. If the HTTP server runs, thus, the configuration is valid. Currently, 
there is no `checkconfig` command.

### Sources

For each iCal (.ics) source, you need:

1. A uniq Source ID
2. The source URI
3. The "Delete" rules to apply

Indeed, the rules are not shared across the sources, allowing you to have multiple sources with different rules.

The Source ID is used in the URL as `http://yourserver:8080/ID`.

When one of the rules listed is met, the event is deleted.

### Rules

The rules are met only when all the criteria are met.

The criteria are targeting a special part of the iCal RFC. The rules may be defined by different form.

For instance, the `title` criteria may be defined as a string, as followed.

```yml
- title: This string is the exact title targeted
```

But it may alsop being defined using a condition, as followed.

```yml
# Exact same thing as previous code
- title:
      condition: exact
      value: This string is the exact title targeted

# It implements also other conditions:
- title:
      condition: not_exact
      value: This string is the title of the only kept event
- title:
      condition: contains
      value: This is all or a part of the targeted event title
- title:
      condition: not_contains
      value: This is all or a part of the kept event title
```

Check the configuration sample, while waiting for proper documentation of each criterion.

## How to add a iCal (ICS) criterion?

First, add the criterion in the [`./internal/rules/`](./internal/rules/) directory. The Golang struct pointer needs 
to implement the `RuleInterface` interface

## How to contribute?

You can find things to do by:

1. (some are easy) Looking the "TODO" in this repo;
2. (easy) Implement a `checkconfig` command;
3. Implement tests;
4. Document code and rules;
5. Implement the event deletion then the event has no ID;
6. Implement what you would like to see implemented;
7. Add some rules.

## Known issues

- The "Edit" rules are not implemented yet.
- The event needs an ID. Some example I saw on internet don't have any ID.
