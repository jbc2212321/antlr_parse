grammar Time;
// Top-level statements

// Tokens
EGT : '>=';
ELT : '<=';
EQ  : '=';
GT  : '>';
LT  : '<';
NUMBER: [0-9]+;
Clause: ('task_end_time'|'task_start_time'|'instance_start_time'|'instance_end_time');
DateType : ('MONTH'|'YEAR'|'WEEK'|'DAY');
YEAR:[1-2][0-1][0-9][0-9];
MONTH:[0][1-9]|[1][0-2];
DAY:[0-3][0-9];
Date: .YEAR.MONTH.DAY.;
TIMEZONECONST:('America/New_York'|'America/Chicago'|'America/Denver'|'America/Los_Angeles'|'America/Vancouver'|'Canada/Eastern'|'Canada/Central'|'Canada/Mountain'|'Canada/Pacific'|'Europe/London'|'Europe/Paris'|'Europe/Berlin'|'Europe/Madrid'|'Europe/Rome'|'Europe/Moscow'|'Asia/Tokyo'|'Asia/Shanghai'|'Asia/Hong_Kong'|'Asia/Seoul'|'Asia/Singapore'|'Asia/Taipei'|'Australia/Sydney'|'Australia/Melbourne'|'Australia/Brisbane'|'Australia/Perth'|'Pacific/Auckland'|'Africa/Johannesburg'|'Brazil/East'|'Mexico/General'|'Antarctica/McMurdo'|'UTC') ;
TIMEZONE:.TIMEZONECONST.;
ToDateTimeOrZero :'toDateTimeOrZero';
DateAdd:'date_add';
WHITESPACE: [ \r\n\t]+ -> skip;

start : expression EOF;
expression
   : expression op=('>'|'<'|'='|'>='|'<=') expression # Compare
   |    Clause # Clause
   |    Date # Date
   |    ToDateTimeOrZero '(' expression ')'                 # ToDateTimeOrZero
   |    ToDateTimeOrZero '(' expression ',' TIMEZONE ')'                 # ToDateTimeOrZeroWithTimeZone
   |    DateAdd '(' DateType ',' '-' NUMBER ',' expression ')'  # DateAdd
   ;
