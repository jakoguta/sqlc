ALTER TABLE venues RENAME TO venue;
ALTER TABLE venue ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();
ALTER TABLE venue DROP COLUMN dropped;
