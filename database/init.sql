
CREATE TABLE users (
  username varchar(64) not null,
  PRIMARY KEY(username)
);

CREATE TABLE assignment (

  assignmentid varchar(64) not null,
  gradingfile varchar(64) not null,
  createdby varchar(64) not null,

  totalscore int,
  PRIMARY KEY(assignmentid),
  FOREIGN KEY(createdby) REFERENCES users(username)
);

CREATE TABLE submission (
  submissionid serial primary key,

  username varchar(64) not null,
  assignment varchar(64) not null,

  isgraded boolean default false,

  obtainedscore int,
  feedback varchar(255),

  FOREIGN KEY(assignment) REFERENCES assignment(assignmentid),
  FOREIGN KEY(username) REFERENCES users(username)

);



