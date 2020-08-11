#include <stdio.h>
#include <time.h>

// retrieve current time
// format as string
// write to standard output

int main()
{
    time_t now;
    struct tm *ts;
    char buf[80];

    // get current time
    now = time(NULL);

    // format and print the time
    // "ddd yyyy-mm-dd hh:mm:ss zzz"
    ts = localtime(&now);
    strftime(buf, sizeof(buf), "%a %Y-%d-%d %H:%M:%S %Z", ts);
    puts(buf);

    return 0;
}
