CREATE TABLE IF NOT EXISTS actors_movies (
    actor_id INTEGER NOT NULL,
    movie_id INTEGER NOT NULL,
    CONSTRAINT fk_actor
        FOREIGN KEY(actor_id) 
        REFERENCES actors(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_movie
        FOREIGN KEY(movie_id) 
        REFERENCES movies(id)
        ON DELETE CASCADE,
    PRIMARY KEY (actor_id, movie_id)
);