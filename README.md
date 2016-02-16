# A&A SMS receiver

This is a very simple receiver for SMS messages using a VOIP line provided by
[Andrews & Arnold ISP](http://aa.net.uk/).

The details of the message posted to the receiver (via HTTP) are documented at
http://aa.net.uk/kb-telecoms-sms.html . A very simple GET endpoint is provided
to retrieve the last message received. A Go client is provided to poll this
endpoint with a timeout.
