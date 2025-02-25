# ICS Filter

Quick and dirty iCal (.ics) filter.

## My use-case

I have third parties who are exporting calendars in a iCal (.ics) address. I add them to my calendar. However, the 
flow is dirty, and they don't allow me to filter event during the export.

By running this tool, I create an HTTP server, acting as a proxy to the third party subscriptions and apply the 
filters each time my calendar wants to refresh the flow.

This server is not made for being blazing fast. Just to have a working server quickly.

## Where is the doc? How to deploy it?

The rest is a work in progress.
