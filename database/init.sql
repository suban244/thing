CREATE TABLE submissions (
  fileid serial primary key,

  username varchar(64) default 'suban',
  filename varchar(64) not null,

  isgraded boolean default false,
  feedback varchar(255)

);
