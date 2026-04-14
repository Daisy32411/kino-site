-- =====================================================
-- CREATE TABLES
-- =====================================================

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    image VARCHAR(500),
    description TEXT,
    year INTEGER,
    director VARCHAR(255),
    file VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS actors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS movie_actors (
    movie_id INTEGER REFERENCES movies(id) ON DELETE CASCADE,
    actor_id INTEGER REFERENCES actors(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, actor_id)
);

CREATE TABLE IF NOT EXISTS movie_genres (
    movie_id INTEGER REFERENCES movies(id) ON DELETE CASCADE,
    genre_id INTEGER REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, genre_id)
);

-- =====================================================
-- INSERT USERS
-- =====================================================

INSERT INTO users (id, email, password) VALUES
(1, 'kamil@gmail.com', '$2a$10$abyJZoAGxtVF9uvZOdeqW.dU.hL2MOV06KisXlPH6uzpSXS/JFcNS'),
(3, 'ka2mil@gmail.com', '$2a$10$qgtAZ1KhDYXB7Js7entxKu8txbcX7Q.Xfhi.sJlS0oiJ7E8HWApN.'),
(4, 'fngjn@gmail.com', '$2a$10$dS1UJdUCXYArViJMlZpYSei08c.xTHqaZoCzetAOYRJK2LV62gXAe'),
(5, 'h@gmail.com', '$2a$10$fIXMQm5vXnr261DDQk.UFeiZp04LIGEoLIWPm2sHo8xX5FqMZxu86')
ON CONFLICT (id) DO UPDATE SET
    email = EXCLUDED.email,
    password = EXCLUDED.password;

-- =====================================================
-- INSERT MOVIES
-- =====================================================

INSERT INTO movies (id, title, slug, image, description, year, director, file) VALUES
(1, 'Django Unchained', 'django-unchained', '/static/images/Django Unchained.jpg', 'Освобождённый раб Джанго объединяется с охотником за головами доктором Кингом Шульцем, чтобы спасти свою жену Брумгильду из рук жестокого плантатора Кэлвина Кэнди. Кровавый вестерн о мести, свободе и справедливости.', 2012, 'Quentin Tarantino', NULL),
(2, 'Inglourious Basterds', 'inglourious-basterds', '/static/images/Inglourious Basterds.jpg', 'Отряд американских солдат-евреев под командованием лейтенанта Альдо Рейна планирует уничтожить нацистское руководство. Параллельно девушка-еврейка Шошанна, владеющая кинотеатром в Париже, готовит собственную месть.', 2009, 'Quentin Tarantino', NULL),
(3, 'Once Upon a Time in Hollywood', 'once-upon-a-time-in-hollywood', '/static/images/Once Upon a Time in Hollywood.jpg', 'Актер-неудачник Рик Далтон и его дублер Клифф Бут пытаются найти своё место в Голливуде конца 60-х, на фоне надвигающихся перемен и трагических событий, изменивших индустрию навсегда.', 2019, 'Quentin Tarantino', NULL),
(4, 'Pulp Fiction', 'pulp-fiction', '/static/images/Pulp Fiction.jpg', 'Переплетающиеся истории двух киллеров, боксёра, жены гангстера и грабителей. Фильм изменил независимый кинематограф 90-х и стал культовым благодаря остроумным диалогам, нелинейному повествованию и незабываемым сценам.', 1994, 'Quentin Tarantino', 'Pulp_Fiction.1994.mp4'),
(5, 'Reservoir Dogs', 'reservoir-dogs', '/static/images/Reservoir Dogs.jpg', 'Шестеро грабителей, использующих кодовые имена, собираются для ограбления ювелирного магазина. После того как операция идет не по плану, оставшиеся в живых пытаются выяснить, кто из них предатель.', 1992, 'Quentin Tarantino', NULL),
(71, 'Interstellar', 'interstellar', '/static/images/Interstellar.jpg', 'Когда засуха, пыльные бури и вымирание растений ставят человечество на грань катастрофы, группа исследователей отправляется через червоточину в другую галактику, чтобы найти планету, пригодную для жизни. Бывший пилот НАСА Купер жертвует всем, чтобы спасти своих детей и будущее человечества.', 2014, 'Christopher Nolan', NULL),
(72, 'Baywatch', 'baywatch', '/static/images/Baywatch.jpg', 'Элитная команда спасателей пляжа Лос-Анджелеса под руководством легендарного Митча Бьюкеннона расследует серию загадочных преступлений на побережье. Вместе с молодым и дерзким новичком они раскрывают криминальную схему, угрожающую их любимому пляжу.', 2017, 'Seth Gordon', NULL),
(73, 'Jurassic World: Fallen Kingdom', 'jurassic-world-fallen-kingdom', '/static/images/Jurassic World Fallen Kingdom.jpg', 'Через три года после разрушения парка «Мир Юрского периода» вулкан на острове Нублар начинает извергаться. Оуэн и Клэр возвращаются на остров, чтобы спасти динозавров от вымирания, но сталкиваются с заговором, который может изменить ход эволюции.', 2018, 'J.A. Bayona', NULL),
(74, 'Black Panther', 'black-panther', '/static/images/Black Panther.jpg', 'После смерти своего отцы принц Т’Чалла возвращается в Ваканду — скрытую, технологически развитую африканскую страну — чтобы занять трон и стать Чёрной Пантерой. Но его правление оспаривает загадочный противник по имени Эрик Киллмонгер.', 2018, 'Ryan Coogler', NULL),
(75, 'Black Panther: Wakanda Forever', 'black-panther-wakanda-forever', '/static/images/Black Panther Wakanda Forever.jpg', 'Королева Рамонда, принцесса Шури и воины Ваканды скорбят о потере Т’Чаллы. Когда скрытая подводная цивилизация Талокан во главе с мутантом Неймором объявляет войну, Ваканда должна объединиться, чтобы защитить своё наследие.', 2022, 'Ryan Coogler', NULL),
(76, 'Mission: Impossible', 'mission-impossible', '/static/images/Mission Impossible 1.jpg', 'Агент Итан Хант, работающий в секретном подразделении «Миссия невыполнима», оказывается обвинённым в предательстве. Чтобы очистить своё имя, он собирает команду и начинает опасную охоту на настоящего крота внутри ЦРУ.', 1996, 'Brian De Palma', NULL),
(77, 'Mission: Impossible 2', 'mission-impossible-2', '/static/images/Mission Impossible 2.jpg', 'Итан Хант получает задание найти и обезвредить опасный вирус «Химера», который украл бывший агент. Для этого ему придётся объединиться с воровкой Ньей, чтобы проникнуть в логово врага и предотвратить глобальную эпидемию.', 2000, 'John Woo', NULL),
(78, 'Mission: Impossible III', 'mission-impossible-3', '/static/images/Mission Impossible 3.jpg', 'Итан Хант покинул оперативную работу, но его втягивает обратно опасный торговец оружием Оуэн Дэвиан. Когда ученицу Итана похищают, он должен спасти её и остановить Дэвиана от продажи смертоносного оружия «Кроличья лапа».', 2006, 'J.J. Abrams', NULL),
(79, 'Mission: Impossible - Ghost Protocol', 'mission-impossible-ghost-protocol', '/static/images/Mission Impossible 4.jpg', 'После теракта в Кремле «Миссия невыполнима» официально ликвидирована. Итан и его команда действуют за пределами закона, используя протокол «Фантом», чтобы остановить безумного стратега, мечтающего о ядерной войне.', 2011, 'Brad Bird', NULL),
(80, 'Mission: Impossible - Rogue Nation', 'mission-impossible-rogue-nation', '/static/images/Mission Impossible 5.jpg', 'Итан Хант сталкивается с Синдикатом — тайной организацией из бывших агентов, работающей на уничтожение спецслужб. В погоне за правдой он встречает загадочную Ильзу Фауст, чьи мотивы остаются неясными.', 2015, 'Christopher McQuarrie', NULL)
ON CONFLICT (id) DO UPDATE SET
    title = EXCLUDED.title,
    slug = EXCLUDED.slug,
    image = EXCLUDED.image,
    description = EXCLUDED.description,
    year = EXCLUDED.year,
    director = EXCLUDED.director,
    file = EXCLUDED.file;

-- =====================================================
-- INSERT ACTORS
-- =====================================================

INSERT INTO actors (id, name) VALUES
(1, 'Jamie Foxx'),
(2, 'Christoph Waltz'),
(3, 'Leonardo DiCaprio'),
(4, 'Samuel L. Jackson'),
(5, 'Kerry Washington'),
(6, 'Brad Pitt'),
(7, 'Mélanie Laurent'),
(8, 'Diane Kruger'),
(9, 'Michael Fassbender'),
(10, 'John Travolta'),
(11, 'Uma Thurman'),
(12, 'Bruce Willis'),
(13, 'Margot Robbie'),
(14, 'Harvey Keitel'),
(15, 'Tim Roth'),
(16, 'Michael Madsen'),
(17, 'Steve Buscemi'),
(239, 'Matthew McConaughey'),
(240, 'Anne Hathaway'),
(241, 'Jessica Chastain'),
(242, 'Michael Caine'),
(243, 'Matt Damon'),
(244, 'Dwayne Johnson'),
(245, 'Zac Efron'),
(246, 'Alexandra Daddario'),
(247, 'Priyanka Chopra'),
(248, 'Chris Pratt'),
(249, 'Bryce Dallas Howard'),
(250, 'Jeff Goldblum'),
(251, 'Rafe Spall'),
(252, 'Chadwick Boseman'),
(253, 'Michael B. Jordan'),
(254, 'Lupita Nyong''o'),
(255, 'Danai Gurira'),
(256, 'Letitia Wright'),
(257, 'Winston Duke'),
(258, 'Tenoch Huerta'),
(259, 'Angela Bassett'),
(260, 'Tom Cruise'),
(261, 'Ving Rhames'),
(262, 'Simon Pegg'),
(263, 'Jeremy Renner'),
(264, 'Rebecca Ferguson'),
(265, 'Jon Voight'),
(266, 'Dougray Scott'),
(267, 'Thandiwe Newton'),
(268, 'Philip Seymour Hoffman'),
(269, 'Paula Patton')
ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name;

-- =====================================================
-- INSERT GENRES
-- =====================================================

INSERT INTO genres (id, name) VALUES
(1, 'Криминал'),
(2, 'Драма'),
(3, 'Боевик'),
(4, 'Комедия'),
(5, 'Вестерн'),
(6, 'Военный'),
(7, 'Фантастика'),
(8, 'Приключения'),
(9, 'Семейный'),
(10, 'Супергероика'),
(11, 'Триллер'),
(12, 'Детектив')
ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name;

-- =====================================================
-- INSERT MOVIE_ACTORS
-- =====================================================

INSERT INTO movie_actors (movie_id, actor_id) VALUES
(71, 239), (71, 240), (71, 241), (71, 242), (71, 243),
(72, 244), (72, 245), (72, 246), (72, 247),
(73, 248), (73, 249), (73, 250), (73, 251),
(74, 252), (74, 253), (74, 254), (74, 255), (74, 256), (74, 257),
(75, 256), (75, 258), (75, 259), (75, 255), (75, 257),
(76, 260), (76, 265), (76, 261),
(77, 260), (77, 266), (77, 267), (77, 261),
(78, 260), (78, 268), (78, 261), (78, 262),
(79, 260), (79, 263), (79, 262), (79, 269),
(80, 260), (80, 264), (80, 262), (80, 261),
(1, 1), (1, 2), (1, 3), (1, 4), (1, 5),
(2, 6), (2, 7), (2, 2), (2, 8), (2, 9),
(3, 3), (3, 6), (3, 13),
(4, 10), (4, 4), (4, 11), (4, 12),
(5, 14), (5, 15), (5, 16), (5, 17)
ON CONFLICT (movie_id, actor_id) DO NOTHING;

-- =====================================================
-- INSERT MOVIE_GENRES
-- =====================================================

INSERT INTO movie_genres (movie_id, genre_id) VALUES
(1, 1), (1, 2), (1, 5),
(2, 1), (2, 2), (2, 6),
(3, 1), (3, 2), (3, 4),
(4, 1), (4, 2), (4, 3),
(5, 1), (5, 2), (5, 3),
(71, 7), (71, 2), (71, 8),
(72, 4), (72, 3), (72, 8),
(73, 7), (73, 3), (73, 8),
(74, 10), (74, 3), (74, 2),
(75, 10), (75, 3), (75, 2),
(76, 3), (76, 11), (76, 12),
(77, 3), (77, 11),
(78, 3), (78, 11), (78, 12),
(79, 3), (79, 11), (79, 8),
(80, 3), (80, 11), (80, 12)
ON CONFLICT (movie_id, genre_id) DO NOTHING;

-- =====================================================
-- UPDATE SEQUENCES
-- =====================================================

SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));
SELECT setval('movies_id_seq', (SELECT MAX(id) FROM movies));
SELECT setval('actors_id_seq', (SELECT MAX(id) FROM actors));
SELECT setval('genres_id_seq', (SELECT MAX(id) FROM genres));