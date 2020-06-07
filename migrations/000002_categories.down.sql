BEGIN;

ALTER TABLE article
DROP COLUMN category_id;

DROP TABLE category;

END;
