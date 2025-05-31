insert into usuarios (nome, nick, email, senha) values ('Usuario 1', 'usuario_1', 'usuario_1@gmail.com', '$2a$10$hyFrprirWtq8.SPZTqBquOWq8Z6o382ub7bb0doq/VZmpU3pLzbW.'),
('Usuario 2', 'usuario_2', 'usuario_2@gmail.com', '$2a$10$hyFrprirWtq8.SPZTqBquOWq8Z6o382ub7bb0doq/VZmpU3pLzbW.'),
('Usuario 3', 'usuario_3', 'usuario_3@gmail.com', '$2a$10$hyFrprirWtq8.SPZTqBquOWq8Z6o382ub7bb0doq/VZmpU3pLzbW.');


insert into seguidores (usuario_id, seguidor_id) values (1, 2), (1, 3), (2, 3);


insert into publicacoes (titulo, conteudo, autor_id) values ('Publicação 1', 'Conteúdo da publicação 1', 1),
('Publicação 2', 'Conteúdo da publicação 2', 1),
('Publicação 3', 'Conteúdo da publicação 3', 2);