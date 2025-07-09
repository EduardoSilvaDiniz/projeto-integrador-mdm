CREATE TABLE associated (
  number_card INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE present (
    associated_id BIGINT NOT NULL,
    meeting_id BIGINT NOT NULL,
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    present BOOLEAN NOT NULL,

    PRIMARY KEY (associated_id, meeting_id),
    FOREIGN KEY (associated_id) REFERENCES associated(id) ON DELETE CASCADE,
    FOREIGN KEY (meeting_id) REFERENCES meeting(id) ON DELETE CASCADE
);
