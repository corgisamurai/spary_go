CREATE DATABASE spary;
CREATE DATABASE spary_test;

grant all on spary.* to spary_admin@localhost identified by 'spary@YRAPS';
grant all on spary_test.* to spary_admin@localhost identified by 'spary@YRAPS';

FLUSH PRIVILEGES;
