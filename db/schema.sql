CREATE TABLE IF NOT EXISTS t_restaurant (
    id char(6) NOT NULL,
    _name text NOT NULL,
    _image text NOT NULL,
    _location point NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS t_product (
    id char(6) NOT NULL,
    restaurant_id char(6) NOT NULL,
    _category text NOT NULL,
    _name text NOT NULL,
    _description text NOT NULL,
    _image text,
    _price_value integer NOT NULL,
    _price_currency text NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO t_restaurant (id, _name, _image, _location)
VALUES (
    '111111',
    'La Bocatería',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    point(41.404564, 2.172751)
);

INSERT INTO t_restaurant (id, _name, _image, _location)
VALUES (
    '222222',
    'Las Hamburguesas',
    'www.laramblabarcelona.com/wp-content/uploads/2017/12/mejores-hamburguesas-de-barcelona-1368x800.jpg',
    point(41.397740, 2.164425)
);
INSERT INTO t_restaurant (id, _name, _image, _location)
VALUES (
    '333333',
    'Poke ball',
    'static.onecms.io/wp-content/uploads/sites/19/2016/08/22/pokeball-poke-bowl-mr.jpg',
    point(41.397208, 2.175534)
);
INSERT INTO t_restaurant (id, _name, _image, _location)
VALUES (
    '444444',
    'Pasti Pasti',
    'www.hogarmania.com/archivos/201401/tipos-de-pasta-xl-668x400x80xX.jpg',
    point(41.423411, 2.187579)
);
INSERT INTO t_restaurant (id, _name, _image, _location)
VALUES (
    '555555',
    'Frankfurt Leoboe',
    'static.wixstatic.com/media/510632_83a9a18e5fd34cd4af839ac36a026473~mv2.png/v1/fill/w_1481,h_1028,al_c/510632_83a9a18e5fd34cd4af839ac36a026473~mv2.png',
    point(41.386885, 2.112450)
);
INSERT INTO t_restaurant (id, _name, _image, _location)
VALUES (
    '666666',
    'Vegan World',
    'media.cnn.com/api/v1/images/stellar/prod/191101102722-vegan-diet-stock.jpg?q=w_3000,h_2002,x_0,y_0,c_fill',
    point(41.387763, 2.169109)
);
INSERT INTO t_restaurant (id, _name, _image, _location)
VALUES (
    '777777',
    'Catalunya Pita House',
    '10619-2.s.cdn12.com/rests/original/101_507074939.jpg',
    point(41.373562, 2.118382)
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00001',
    '111111',
    'preferidos',
    'Bocata de atún',
    'blablablablablablalalbal vbakb kvbadk',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    300,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00002',
    '111111',
    'bebidas',
    'Cocacola',
    'muy fría',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    150,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00003',
    '111111',
    'veganos',
    'Bocadillo de aguacate',
    'muy vegano',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    500,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00004',
    '222222',
    'clásicos',
    'Hamburguesa completa',
    'con todo',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    900,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00005',
    '222222',
    'clásicos',
    'Hamburguesa doble',
    'doble hamburguesa',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    1100,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00006',
    '333333',
    'clásicos',
    'poke bowl de taún',
    'Arroz, tomate, canónigos, atún',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    700,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00007',
    '333333',
    'clásicos',
    'poke bowl de taún',
    'Arroz, tomate, canónigos, atún, salsa de soja',
    'static.onecms.io/wp-content/uploads/sites/19/2016/08/22/pokeball-poke-bowl-mr.jpg',
    700,
    'EUR'
);
INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00008',
    '333333',
    'clásicos',
    'poke bowl de salmón',
    'Arroz, piña, lechuga, salmón, salsa teriyaki',
    'static.onecms.io/wp-content/uploads/sites/19/2016/08/22/pokeball-poke-bowl-mr.jpg',
    700,
    'EUR'
);
INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00009',
    '333333',
    'pokemons',
    'Magikarp',
    'Quinoa, piña, lechuga, salmón, salsa teriyaki',
    'static.onecms.io/wp-content/uploads/sites/19/2016/08/22/pokeball-poke-bowl-mr.jpg',
    700,
    'EUR'
);
INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00010',
    '333333',
    'pokemons',
    'Pidgey',
    'Quinoa, tomate, canónigos, pollo rebozado, salsa teriyaki',
    'static.onecms.io/wp-content/uploads/sites/19/2016/08/22/pokeball-poke-bowl-mr.jpg',
    700,
    'EUR'
);
INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00011',
    '333333',
    'pokemons',
    'Pickachu',
    'Arroz, tomate cherry, canónigos, pollo rebozado, salsa soja',
    'static.onecms.io/wp-content/uploads/sites/19/2016/08/22/pokeball-poke-bowl-mr.jpg',
    700,
    'EUR'
);
INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00012',
    '333333',
    'pokemons',
    'Charmander',
    'Quinoa, piña, canónigos, salmón flameado, salsa picante',
    'static.onecms.io/wp-content/uploads/sites/19/2016/08/22/pokeball-poke-bowl-mr.jpg',
    700,
    'EUR'
);
INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00013',
    '333333',
    'pokemons',
    'Bulbasaur',
    'Quinoa, piña, alcachofa, atún flameado, salsa soja',
    'static.onecms.io/wp-content/uploads/sites/19/2016/08/22/pokeball-poke-bowl-mr.jpg',
    700,
    'EUR'
);
INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00014',
    '333333',
    'pokemons',
    'Bulbasaur',
    'Arroz, mango, lechuga, heura, sesamo',
    'static.onecms.io/wp-content/uploads/sites/19/2016/08/22/pokeball-poke-bowl-mr.jpg',
    700,
    'EUR'
);
