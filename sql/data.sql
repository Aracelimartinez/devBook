-- USERS
INSERT INTO users (name, nick, email, password)
VALUES
('Araceli Martinez', 'amv94', 'email@example.com', '$2a$10$XclUh7TJAdRu/TCyHcdyDe1brzPBiZUQUjv6fv8CMUg2IJK4H3axS'),
('Diego Pico', 'dmd94', 'email2@example.com', '$2a$10$XclUh7TJAdRu/TCyHcdyDe1brzPBiZUQUjv6fv8CMUg2IJK4H3axS'),
('Usuario 3', 'usr3', 'email3@example.com', '$2a$10$XclUh7TJAdRu/TCyHcdyDe1brzPBiZUQUjv6fv8CMUg2IJK4H3axS'),
('Usuario 4', 'usr4', 'email4@example.com', '$2a$10$XclUh7TJAdRu/TCyHcdyDe1brzPBiZUQUjv6fv8CMUg2IJK4H3axS'),
('Usuario 5', 'usr5', 'email5@example.com', '$2a$10$XclUh7TJAdRu/TCyHcdyDe1brzPBiZUQUjv6fv8CMUg2IJK4H3axS');

-- FOLLOWERS
INSERT INTO followers(user_id, follower_id)
VALUES
(1,2),
(2,1),
(1,3),
(3,1),
(1,4),
(4,1),
(1,5),
(5,1),
(3,2),
(5,4),
(4,3),
(2,5);

-- POSTS
-- Publicación para Araceli Martinez (user_id = 1)
INSERT INTO posts (title, content, author_id, likes)
VALUES
('Publicación de Araceli', 'Contenido de la publicación de Araceli', 1, 10),
('Publicación de Diego', 'Contenido de la publicación de Diego', 2, 5),
('Publicación de Usuario 3', 'Contenido de la publicación de Usuario 3', 3, 8),
('Publicación de Usuario 4', 'Contenido de la publicación de Usuario 4', 4, 12),
('Publicación de Usuario 5', 'Contenido de la publicación de Usuario 5', 5, 3);
