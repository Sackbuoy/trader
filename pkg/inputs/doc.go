/*
This package contains the part of the pipeline where we gather a list of
equities that we want to evaluate with our algorithm. For example:
  - All equities available on whatever platform we happen to be using to get data
  - All equities available on the NASDAQ, or NYSE
  - Equities stored in the local filesystem(useful for testing)
  - Pull equities that are in a certain field: e.g. Tech stocks, or maybe just
    ETF's in general, or stuff like Emerging Market equities.

This step is not intended to do any evaluation of the stocks at all, it solely
grabs them based on information that never really changes
*/
package inputs
