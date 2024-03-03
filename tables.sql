USE s3db;

CREATE TABLE IF NOT EXISTS buckets (
  id bigint NOT NULL,
  name varchar(255) NOT NULL,
  public boolean NOT NULL DEFAULT false,
  description text,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS files (
  id bigint NOT NULL,
  bucket_id bigint NOT NULL,
  name varchar(255) NOT NULL UNIQUE,
  size bigint NOT NULL,
  mime_type varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT files_ibfk_1 FOREIGN KEY (bucket_id) REFERENCES buckets (id) ON DELETE CASCADE ON UPDATE CASCADE
);
