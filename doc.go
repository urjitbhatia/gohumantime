/*

Package gohumantime is a Human readable time interval parser for golang

    import "github.com/urjitbhatia/gohumantime"

    ...

    ToMilliseconds("ten minutes and four seconds")
    ToMilliseconds('one minute')
    ToMilliseconds('2.3 minutes')
    ToMilliseconds('3 days, 4 hours and 36 seconds')

Supported Units

The following time unit words are supported

  - `seconds`
  - `minutes`
  - `hours`
  - `days`
  - `weeks`
  - `months` -- assumes 30 days
  - `years` -- assumes 365 days

Numeric word representations

gohumantime supports numbers up to ten being written out in English.
Beyond then, it understands:

  `fifteen`
and then every tens unit from 
  'twenty', 'thirty'... onto 'hundred'

*/
package gohumantime
