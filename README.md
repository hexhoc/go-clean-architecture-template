# GO Clean architecture template

## Project structure
1. /cmd/app

The main (main) applications for this project.

The directory name for each application corresponds to the name of the executable you want to get (eg /cmd/myapp).

2. /internal

Private application and library code. This is code that you don't want others to import into their applications or libraries.

3. /pkg

Library code that can be used by external applications (eg /pkg/mypubliclib). Other projects will import these libraries expecting them to work.

