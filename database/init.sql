CREATE TABLE submissions (
  fileid serial primary key,

  username varchar(64) not null,
  filename varchar(64) not null,

  isgraded boolean default false,
  obtainedscore int,
  maxscore int,

  feedback varchar(255)
);
