create table rates (
    id          integer constraint rates_pk primary key autoincrement,
    code        varchar(255),
    code_in     varchar(255),
    name        varchar(255),
    high        varchar(255),
    low         varchar(255),
    var_bid     varchar(255),
    pct_change  varchar(255),
    bid         varchar(255),
    ask         varchar(255),
    timestamp   varchar(255),
    create_date varchar(255)
);
