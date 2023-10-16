CREATE TABLE submissions (
  fileID varchar(64) primary key,

  username varchar(64) not null,
  filename varchar(64) not null,

  isGraded boolean default false

);
