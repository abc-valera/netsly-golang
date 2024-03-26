-- Here lies the fixtures for the database.
--
-- There are:
--  5 users
--  10 jokes
--  20 likes
--  10 comments
--  5 rooms
--  10 room members
--  10 room messages
--
-- The data is inserted in the order of the tables' dependencies.
-- 
-- 
--
INSERT INTO
  "User" (
    id,
    username,
    email,
    hashed_password,
    fullname,
    status,
    created_at
  )
VALUES
  (
    '00000000-0000-0000-0000-000000000000',
    'abc-valera',
    'abc-valera@gmail.com',
    'hashed_12sagas1sag',
    'Valeriy',
    'active',
    NOW ()
  ),
  (
    '11111111-1111-1111-1111-111111111111',
    'johndoe',
    'johndoe@gmail.com',
    'hashed_12y4721rui',
    'John Doe',
    'active',
    NOW ()
  ),
  (
    '22222222-2222-2222-2222-222222222222',
    'janedoe',
    'janedoe@gmail.com',
    'hashed_18asxvnio78s',
    'Jane Doe',
    'inactive',
    NOW ()
  ),
  (
    '33333333-3333-3333-3333-333333333333',
    'tomjones',
    'tomjones@gmail.com',
    'hashed_34y4721rui',
    'Tom Jones',
    'active',
    NOW ()
  ),
  (
    '44444444-4444-4444-4444-444444444444',
    'marysmith',
    'marysmith@gmail.com',
    'hashed_45asxvnio78s',
    'Mary Smith',
    'inactive',
    NOW ()
  );

INSERT INTO
  "Joke" (id, title, text, explanation, created_at, user_id)
VALUES
  (
    '00000000-0000-0000-0000-000000000000',
    'Muddy dirt',
    'What did the dirt say to the rain? If you keep this up, my name will be mud!',
    'The phrase ''my name will be mud'' is a pun on the word ''mud'' and the phrase ''my name will be mud'' which means ''I will be in trouble''.',
    NOW (),
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    '11111111-1111-1111-1111-111111111111',
    'Light bulb',
    'How many programmers does it take to change a light bulb? None, that''s a hardware problem!',
    'This joke is a play on the division of responsibilities in computer science. Programmers typically handle software problems, not hardware ones.',
    NOW (),
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    '22222222-2222-2222-2222-222222222222',
    'Java',
    'Why don''t programmers like nature? It has too many bugs!',
    'In programming, a ''bug'' is an error or flaw in the code.',
    NOW (),
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    '33333333-3333-3333-3333-333333333333',
    'Python',
    'Why do Python programmers prefer snakes? Because they''re not Java!',
    'This joke is a play on the names of two programming languages, Java and Python.',
    NOW (),
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    '44444444-4444-4444-4444-444444444444',
    'C++',
    'Why do C++ programmers need glasses? Because they can''t C#!',
    'This joke is a play on the names of two programming languages, C++ and C# (pronounced ''C Sharp'').',
    NOW (),
    '33333333-3333-3333-3333-333333333333'
  ),
  (
    '55555555-5555-5555-5555-555555555555',
    'JavaScript',
    'Why did the JavaScript file go to the therapist? Because it had too many unresolved promises!',
    'In JavaScript, a ''Promise'' is an object that may produce a single value some time in the future.',
    NOW (),
    '44444444-4444-4444-4444-444444444444'
  ),
  (
    '66666666-6666-6666-6666-666666666666',
    'SQL',
    'Why don''t databases make good secret keepers? Because they can''t keep a query to themselves!',
    'A ''query'' is a request for information from a database.',
    NOW (),
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    '77777777-7777-7777-7777-777777777777',
    'HTML',
    'Why did the HTML go to school? Because it wanted to learn the <head> from the <body>!',
    'In HTML, the <head> and <body> are two main parts of a webpage.',
    NOW (),
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    '88888888-8888-8888-8888-888888888888',
    'CSS',
    'Why did the CSS file go to jail? Because it had too many !important rules!',
    'In CSS, ''!important'' is a keyword added to a rule to make it more important than other rules.',
    NOW (),
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    '99999999-9999-9999-9999-999999999999',
    'React',
    'Why did the React component feel relieved? Because it was in a state of suspense!',
    'In React, a ''state'' is an object that holds some information that may change over the lifetime of the component.',
    NOW (),
    '33333333-3333-3333-3333-333333333333'
  );

INSERT INTO
  "Like" (created_at, user_id, joke_id)
VALUES
  (
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '33333333-3333-3333-3333-333333333333'
  ),
  (
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    NOW (),
    '44444444-4444-4444-4444-444444444444',
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '44444444-4444-4444-4444-444444444444'
  ),
  (
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '33333333-3333-3333-3333-333333333333'
  ),
  (
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '44444444-4444-4444-4444-444444444444'
  ),
  (
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    NOW (),
    '44444444-4444-4444-4444-444444444444',
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '55555555-5555-5555-5555-555555555555'
  ),
  (
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '66666666-6666-6666-6666-666666666666'
  ),
  (
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '77777777-7777-7777-7777-777777777777'
  ),
  (
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '88888888-8888-8888-8888-888888888888'
  ),
  (
    NOW (),
    '44444444-4444-4444-4444-444444444444',
    '99999999-9999-9999-9999-999999999999'
  ),
  (
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '66666666-6666-6666-6666-666666666666'
  ),
  (
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '77777777-7777-7777-7777-777777777777'
  ),
  (
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '88888888-8888-8888-8888-888888888888'
  ),
  (
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '99999999-9999-9999-9999-999999999999'
  ),
  (
    NOW (),
    '44444444-4444-4444-4444-444444444444',
    '55555555-5555-5555-5555-555555555555'
  );

INSERT INTO
  "Comment" (id, text, created_at, user_id, joke_id)
VALUES
  (
    '00000000-0000-0000-0000-000000000000',
    'This joke is hilarious!',
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '88888888-8888-8888-8888-888888888888'
  ),
  (
    '11111111-1111-1111-1111-111111111111',
    'What a joke! Very funny',
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    '22222222-2222-2222-2222-222222222222',
    'That''s a good one! Made my day.',
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    '33333333-3333-3333-3333-333333333333',
    'I can''t stop laughing at this!',
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    '44444444-4444-4444-4444-444444444444',
    'This joke is so relatable!',
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    '55555555-5555-5555-5555-555555555555',
    'I didn''t get it at first, but now I can''t stop laughing!',
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '33333333-3333-3333-3333-333333333333'
  ),
  (
    '66666666-6666-6666-6666-666666666666',
    'This joke made my day!',
    NOW (),
    '44444444-4444-4444-4444-444444444444',
    '44444444-4444-4444-4444-444444444444'
  ),
  (
    '77777777-7777-7777-7777-777777777777',
    'I shared this joke with my team, they loved it!',
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '55555555-5555-5555-5555-555555555555'
  ),
  (
    '88888888-8888-8888-8888-888888888888',
    'This is the funniest joke I''ve heard in a while!',
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '66666666-6666-6666-6666-666666666666'
  ),
  (
    '99999999-9999-9999-9999-999999999999',
    'I can''t stop laughing at this joke!',
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '77777777-7777-7777-7777-777777777777'
  );

INSERT INTO
  "Room" (id, name, description, created_at, creator_id)
VALUES
  (
    '00000000-0000-0000-0000-000000000000',
    'Dad jokes',
    'A room for sharing dad jokes! :D',
    NOW (),
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    '11111111-1111-1111-1111-111111111111',
    'Programming Humor',
    'A room for sharing programming jokes and humor! :D',
    NOW (),
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    '22222222-2222-2222-2222-222222222222',
    'Fishing jokes',
    'Fishers from all over the world sharing their best ones',
    NOW (),
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    '33333333-3333-3333-3333-333333333333',
    'Gaming memes',
    'A room for sharing gaming memes and jokes',
    NOW (),
    '33333333-3333-3333-3333-333333333333'
  ),
  (
    '44444444-4444-4444-4444-444444444444',
    'AI Chat',
    'A room for discussing artificial intelligence and machine learning! :D',
    NOW (),
    '44444444-4444-4444-4444-444444444444'
  );

INSERT INTO
  "RoomMember" (created_at, user_id, room_id)
VALUES
  (
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '33333333-3333-3333-3333-333333333333'
  ),
  (
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '44444444-4444-4444-4444-444444444444'
  ),
  (
    NOW (),
    '44444444-4444-4444-4444-444444444444',
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    NOW (),
    '44444444-4444-4444-4444-444444444444',
    '33333333-3333-3333-3333-333333333333'
  );

INSERT INTO
  "RoomMessage" (id, text, created_at, user_id, room_id)
VALUES
  (
    '00000000-0000-0000-0000-000000000000',
    'Hi! How are you?',
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '00000000-0000-0000-0000-000000000000'
  ),
  (
    '11111111-1111-1111-1111-111111111111',
    'Hello, everyone!',
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    '22222222-2222-2222-2222-222222222222',
    'Good morning, folks!',
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    '33333333-3333-3333-3333-333333333333',
    'What''s up, gamers?',
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '33333333-3333-3333-3333-333333333333'
  ),
  (
    '44444444-4444-4444-4444-444444444444',
    'Hi, AI enthusiasts!',
    NOW (),
    '44444444-4444-4444-4444-444444444444',
    '44444444-4444-4444-4444-444444444444'
  ),
  (
    '55555555-5555-5555-5555-555555555555',
    'Hey, how''s it going?',
    NOW (),
    '00000000-0000-0000-0000-000000000000',
    '11111111-1111-1111-1111-111111111111'
  ),
  (
    '66666666-6666-6666-6666-666666666666',
    'Good evening, everyone!',
    NOW (),
    '11111111-1111-1111-1111-111111111111',
    '22222222-2222-2222-2222-222222222222'
  ),
  (
    '77777777-7777-7777-7777-777777777777',
    'How''s everyone doing?',
    NOW (),
    '22222222-2222-2222-2222-222222222222',
    '33333333-3333-3333-3333-333333333333'
  ),
  (
    '88888888-8888-8888-8888-888888888888',
    'Any new games out there?',
    NOW (),
    '33333333-3333-3333-3333-333333333333',
    '44444444-4444-4444-4444-444444444444'
  ),
  (
    '99999999-9999-9999-9999-999999999999',
    'What''s the latest in AI?',
    NOW (),
    '44444444-4444-4444-4444-444444444444',
    '00000000-0000-0000-0000-000000000000'
  );