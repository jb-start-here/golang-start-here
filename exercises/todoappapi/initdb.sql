CREATE TABLE
  public.todos (
    id serial NOT NULL,
    description text NOT NULL,
    done boolean NOT NULL,
    duedate date NULL
  );

ALTER TABLE
  public.todos
ADD
  CONSTRAINT todos_pkey PRIMARY KEY (id);

INSERT INTO "todos" ("description", "done", "duedate", "id") VALUES ('pet dog', true, '2022-12-25', 1);
INSERT INTO "todos" ("description", "done", "duedate", "id") VALUES ('solve a murder mystery', false, '2023-11-25', 2);
