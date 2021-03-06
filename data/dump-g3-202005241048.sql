PGDMP     
    0    
            x            g3    9.5.1    11.2 !    l           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            m           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            n           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                       false            o           1262    26138    g3    DATABASE     �   CREATE DATABASE g3 WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'Russian_Russia.1251' LC_CTYPE = 'Russian_Russia.1251';
    DROP DATABASE g3;
             postgres    false                        2615    26139    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
             postgres    false            p           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                  postgres    false    7            �            1259    26140    map    TABLE     B   CREATE TABLE public.map (
    id text NOT NULL,
    chunk text
);
    DROP TABLE public.map;
       public         postgres    false    7            �            1259    26146    masterships    TABLE     �   CREATE TABLE public.masterships (
    id integer NOT NULL,
    name text,
    short_name text,
    min integer,
    max integer
);
    DROP TABLE public.masterships;
       public         postgres    false    7            �            1259    26152    mastery_items    TABLE     +  CREATE TABLE public.mastery_items (
    id integer NOT NULL,
    mastery text,
    category text,
    ingredient text,
    name text,
    rarity integer,
    areas text,
    area_size integer,
    min integer,
    max integer,
    is_countable boolean,
    is_liquid boolean,
    "limit" integer
);
 !   DROP TABLE public.mastery_items;
       public         postgres    false    7            �            1259    26158    params    TABLE     =   CREATE TABLE public.params (
    key text,
    value text
);
    DROP TABLE public.params;
       public         postgres    false    7            �            1259    26164    person_inventory    TABLE     �   CREATE TABLE public.person_inventory (
    id text NOT NULL,
    person_id integer NOT NULL,
    item_id integer NOT NULL,
    weight text,
    quality text,
    creation_date integer,
    exp_date integer,
    is_deleted boolean
);
 $   DROP TABLE public.person_inventory;
       public         postgres    false    7            �            1259    26170    person_masterships    TABLE     |   CREATE TABLE public.person_masterships (
    person_id integer NOT NULL,
    mastery_id integer NOT NULL,
    skill text
);
 &   DROP TABLE public.person_masterships;
       public         postgres    false    7            �            1259    26176    persons    TABLE     �   CREATE TABLE public.persons (
    id integer NOT NULL,
    name text,
    age integer,
    is_male boolean,
    chunk_id text,
    day_action text
);
    DROP TABLE public.persons;
       public         postgres    false    7            �            1259    26214    t_user    TABLE     �   CREATE TABLE public.t_user (
    user_id bigint NOT NULL,
    username character varying,
    first_name character varying,
    last_name character varying,
    lang character varying,
    creation_date timestamp without time zone NOT NULL
);
    DROP TABLE public.t_user;
       public         postgres    false    7            b          0    26140    map 
   TABLE DATA               (   COPY public.map (id, chunk) FROM stdin;
    public       postgres    false    181            c          0    26146    masterships 
   TABLE DATA               E   COPY public.masterships (id, name, short_name, min, max) FROM stdin;
    public       postgres    false    182            d          0    26152    mastery_items 
   TABLE DATA               �   COPY public.mastery_items (id, mastery, category, ingredient, name, rarity, areas, area_size, min, max, is_countable, is_liquid, "limit") FROM stdin;
    public       postgres    false    183            e          0    26158    params 
   TABLE DATA               ,   COPY public.params (key, value) FROM stdin;
    public       postgres    false    184            f          0    26164    person_inventory 
   TABLE DATA               x   COPY public.person_inventory (id, person_id, item_id, weight, quality, creation_date, exp_date, is_deleted) FROM stdin;
    public       postgres    false    185            g          0    26170    person_masterships 
   TABLE DATA               J   COPY public.person_masterships (person_id, mastery_id, skill) FROM stdin;
    public       postgres    false    186            h          0    26176    persons 
   TABLE DATA               O   COPY public.persons (id, name, age, is_male, chunk_id, day_action) FROM stdin;
    public       postgres    false    187            i          0    26214    t_user 
   TABLE DATA               _   COPY public.t_user (user_id, username, first_name, last_name, lang, creation_date) FROM stdin;
    public       postgres    false    188            �           2606    26183    map map_id_key 
   CONSTRAINT     G   ALTER TABLE ONLY public.map
    ADD CONSTRAINT map_id_key UNIQUE (id);
 8   ALTER TABLE ONLY public.map DROP CONSTRAINT map_id_key;
       public         postgres    false    181            �           2606    26185    masterships masterships_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.masterships
    ADD CONSTRAINT masterships_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.masterships DROP CONSTRAINT masterships_pkey;
       public         postgres    false    182            �           2606    26187     mastery_items mastery_items_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.mastery_items
    ADD CONSTRAINT mastery_items_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.mastery_items DROP CONSTRAINT mastery_items_pkey;
       public         postgres    false    183            �           2606    26189    params params_key_key 
   CONSTRAINT     O   ALTER TABLE ONLY public.params
    ADD CONSTRAINT params_key_key UNIQUE (key);
 ?   ALTER TABLE ONLY public.params DROP CONSTRAINT params_key_key;
       public         postgres    false    184            �           2606    26191 (   person_inventory person_inventory_id_key 
   CONSTRAINT     a   ALTER TABLE ONLY public.person_inventory
    ADD CONSTRAINT person_inventory_id_key UNIQUE (id);
 R   ALTER TABLE ONLY public.person_inventory DROP CONSTRAINT person_inventory_id_key;
       public         postgres    false    185            �           2606    26193    persons persons_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.persons
    ADD CONSTRAINT persons_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.persons DROP CONSTRAINT persons_pkey;
       public         postgres    false    187            �           1259    26220    t_user_user_id_idx    INDEX     O   CREATE UNIQUE INDEX t_user_user_id_idx ON public.t_user USING btree (user_id);
 &   DROP INDEX public.t_user_user_id_idx;
       public         postgres    false    188            �           2606    26194 $   person_inventory person_inventory_fk    FK CONSTRAINT     �   ALTER TABLE ONLY public.person_inventory
    ADD CONSTRAINT person_inventory_fk FOREIGN KEY (person_id) REFERENCES public.persons(id);
 N   ALTER TABLE ONLY public.person_inventory DROP CONSTRAINT person_inventory_fk;
       public       postgres    false    2026    185    187            �           2606    26199 &   person_inventory person_inventory_fk_1    FK CONSTRAINT     �   ALTER TABLE ONLY public.person_inventory
    ADD CONSTRAINT person_inventory_fk_1 FOREIGN KEY (item_id) REFERENCES public.mastery_items(id);
 P   ALTER TABLE ONLY public.person_inventory DROP CONSTRAINT person_inventory_fk_1;
       public       postgres    false    185    2020    183            �           2606    26204 (   person_masterships person_masterships_fk    FK CONSTRAINT     �   ALTER TABLE ONLY public.person_masterships
    ADD CONSTRAINT person_masterships_fk FOREIGN KEY (person_id) REFERENCES public.persons(id);
 R   ALTER TABLE ONLY public.person_masterships DROP CONSTRAINT person_masterships_fk;
       public       postgres    false    186    2026    187            �           2606    26209 *   person_masterships person_masterships_fk_1    FK CONSTRAINT     �   ALTER TABLE ONLY public.person_masterships
    ADD CONSTRAINT person_masterships_fk_1 FOREIGN KEY (mastery_id) REFERENCES public.masterships(id);
 T   ALTER TABLE ONLY public.person_masterships DROP CONSTRAINT person_masterships_fk_1;
       public       postgres    false    186    182    2018            b     x��W]�\5}f�}#'q��(/<Е
�>8�#v�j�PA��O�s��mF�/�;ʙ�����m@"S��ހ| HE?��ZoUb�����W���'��f�~9�����f��ݎ���,�?�^.�|!�?�z'7/�t��~7����J�]�Z����f�n�������6�׻�+�?S�ŵ><ч/v������.��o�\Ab��Zd��[��JE[r���`�B�eh�/o�Ė��޺���V�5P4S�59�6�bK'������c+�%��
uH]_Y8���v��OB�:�� �0��J����"5A�vU�v�i�(�BMН�@9jfb�b���W���v���Y|�����z��ܣ-z�\ pU�%�z�!@��G����3�;���p�J�� T��\��Z0�F�]�9p�����ý���(�p�tC� Es]��"��$/��ʜ������{Q�-�]8{-4*U���Ј3i�]�x�8�(��Pt�h7{�J���-#�>�T��l\]�8����Q�����?�9�݊~z@���J�z0�,�5n�5X
.��ȇL�`����#��n�_�2�9F���հ�4*i��bq�Itd�f�w�����K��j=�R+�M�i?$��a\=l;(t��V��(��!�D����N�.�N%ť���	��\��d����]���n�'��X�jKp���f��&?u�4-���ѣ���u��z�A�x����v0�ӌTX0;5eh#Pu�b�b�%�i���`����b��CCH��-M3�]%�Si5���vP����_�8��K�Y����D��֡dn�
v�s�7�c�fq|g?.�'�׭q�rn��i���:�Ć�f����k������Bs���V���={pdL�R��k��ij3��4��>k)�[kf;��iF�ht��u\!�h,
I��J(�I�1�(�e��y��|]k�@R0�~(Y�3�ac;�GS�
���]�O��#S��Ե&�u>��<Nse?}~vv�73�      c   �   x�E�K
�0E�/��
����`5N��W�����i! ���Q����^�1��_X���NcG����ec˪ C&�TL���=f���=��0����C"=�⿹Z46�����6��u��ť���;��Re�LQ      d   <  x����NA�?o��hL�Vʻ��b���J��E�!DM +���t��������3����N��4l	��s��K]g�ުշ��o�Y��B�ag��w��凎�4믫����*޸��x�����x�@8el��1��)�7 ��{�{�Ǻy�p
��}6Ώ�K#M9�`cD2a�ኧ���+e�����|�e��%%��L�5�����Ǭ�/��خl�?H�|[�d0�
�Gl��_��k \�X����A(EtM:��(w@�xJ0	�T4���M:��}���C(+ ��� 艆rD̓I��I��P/�*�$�sQ���@pM2��
"T�4Wa�c��#����g�I����w��ɓ�;�v�ګ��0ج�����1�QHa�Ѭ��E��Zc[�2!��'(x����J�`̏�Թ�b����?JyUT�o��#��
1��q��M9��s�(�PZ��0��������WHn^�.�8P��)�i�����>� �q	9����
1[۫�/���^�K�,��vdښ�B�ao@z)�bG���@@��s��S�-wR$=���1�_�ÆԜBC{�(`�
�o��K�?^'����zvwo��>'��$�-��&�|����^"�=Sw�x,�P<��� �?d-z���:������j@k)�(eo"<$�iŎ�Vxt�-1�`���jc�����G*]�/�5C�G�.w`o*�\��w4;9Q��ړ��Z�I�j��F��mH+G�9�͒�Mq).�t����B>�_���[��7�%�P��ާ�{V��%Jg\q.�P�E`ek�rHR�3SG�,[8?h�$1�FN���,nk:ip�*´K��BZ��:PĈv�@LU r9a�����Aс�]I�����	��$��ľ��Uk6�*��M��
�L7�Լ�մ�c��7{���=H˩a5���:hFtȥ�RSF��^G��?`���;b>!r�fCM�:�� �n&N��_G��>G��sA��uP�M���įc.�������5��kv��G�P��u����tA�S\�e���
�9��      e      x�KI,I崴�4�����  �5      f   �  x���k�4+���Z`����+��k^"��Dd�zt�R��y�`|W�5ͬ����iIE��Һw��r����h�2�(������S6�?���z��t`��1�b�)����t��oŘ�K3��,��+n�C��=k�J�4������Zj��GZQ{@���$%�!5��r�M�歔صM��$��d��U�Zh��	i�B�q,��2�e��,>��V!RPL�4T��C�Ԛ�G�%�lԦ1�vв�o-���2���ʶ]X�G̾����{���_h|ſg�A�7�/�y�#긤#Z�Q�����QY��Ɂ���hг�������j�Qz���^��5�]���R��Q������B��8;��Q9y���7�a�2ؼ��˦�^l���T��\�O�4��|�=���8L��L���&F�G�b��'�c_��Q��/C�m���n�2]L�����NN���z�Ά�[v���w	��N���!�3E�d����SH�����+t���t�Kfy��mo�)E>�{���S~�N�nK�!<�cx,w�������w�N�t�VT��68���[�E<D�kV�Y뾢���n�����(jC�&�f)���4y��OS��s��xl�q*\
�|h�*���d��A�\`��8�Y����)�ė�NC��u�^x=@��*�ܠs�ҋI�Oi���N�1n��f�d��;�����˱���`�+J�ee�u��`�/��B�6.Зl;	q�&���'���5ӧ1L(��<�V�ֺ�_Rs�q���C�Z�E���㦤�kEW���m�t@��nH��ﰾ��D�;ܓ+z�l�Q��i�'g������D#�W��z�.��؉f����Ȱ?b����zG�!:E%^�{�N�j�Jw�����:Y|\"w���+<�:�Qr�Lڳ�	�VX�w\�JQ�X��4G���B�v*�B����zN�x�f[�'���(��N\U��XV�����m	z��d��2�[��L�]q଄+B#�*����L�qb��K�X��C��%&hպ4�SF�m�
uC	�
� �:E?�ZK�$cQ�9\�0:ݤ��l��,���
Te�(�]�dCC08�ԊF(�T�z.�(ָwfӕ�:K�����H��dU��(:-JG�K�k�DG4�l�%�����!_�ޤԘ���}�R�,��/��
�Q��Aj�+�\��9SS�� ^+�٩h��.!�e}8���f�]�P��<4\=4��r:Xo��&�!�LOM���ry�m%�%�&��F2=�ϊJ ׀6��u��`T"�5:�)��kyB�;����%�qz�J,�d��G͢b����g4�5D�:TlYCŨj[�uSp�\=�iEa�����KP��U�x"1Ѥf��dE��RNō�5<��C*t���ǩd9p�3�݁qi�kU�C*<(�OĖ�l��AÛT&6�`�y|:uZn8uvQK.83����a��{BXBbe��KE�C=P΍�9t�ꌍ��l���J�q�8�R�z������M���sR�MS٘��u�J�𣸠'�S��('�ם�5e�3�3	B�|&n�k�$�^�d�����+Y8��\�ش�{�m��V�PTf�X\[N)�v��l(j�>�s�!	� ����� ���+����n�%���.`���:�1��:���:[����� ,e��*e���igﭡlT&�XBA��9�Lj��F�-N��@G�ZQD6S�"�Z�{'D��E٠K�L>�݊2TQ�U��p%m���i1����w�,?i�8(�³�-�K��4���R���ez��:�,�.�M%��<����xEsk�u�UDd GV���":�V��c�X�T�e�T������Z��G︭�x԰��So?�[�؎ǵ>���Z�k��{�櫡�Tmq$0
D�h2��[L�#��y��滛g�����t9�[�֦�-'=����y�릦;��3#=D�~}�������NÉ��ܠ'#���-`���hA9��8��S�����y�
�Դ��p�N޾�ߡS˳��ߠ'ݳy(uje'���N�},v6'������ϯ;5����x뎶X/6�8>��ڦzWQܠ��w�����=�:�ݨ�����iJ�M�o����:��w����{��d,;�磆i�I�o��7kX��N�����ew����Wh����o��w���=�H=1������	:/$�m u�N�<>E�c�cO~�����9}lps�����s#��.�ߠS�8�p5!��fԓ��5��g�y����U�z1Τ蜟�E��/u/7�y1��=���A���{��(��N<9�S������)���R�:�ޥ�����;yڄN#Dk9�o.�ѸG�����G�ܶI���@2�Pk誂C��b����h�N�;[o�ťC�u�I�X����(��&F��~��J7]8ze�E�ͦ�I�rH�:���8K\Wd�I�W��>_��먊6�R�����i���Y�*I��o�1����%9�������bn������E+��Bl@�l��R�+�8M��kɑZ�T���dL�O��y����9�kO��w߷��m��ba���R,UJ�����g���3�)�Oʬ×��I6vы�OЋ��N>�>U�Ŧ�Q�i\��,�pm_��e	�k��=]\�7�I����ܠ��"oГ��Ou2E�����l��t���w�(�x*d��ߡ���=�\��zѐ=�:�6w���}��8��t�����Ks�ዽ�零�=�v^����O�n�n�ߧ��=�w�>J�HcGJ�d"�� *f4�l��9��Bێ���|�AS�
��� rq(�9˸L)�Ě�\ڳ%�@��z��,V����^�J�{2�ykUV
���N�ŏ�����ig��E֑t.q|��I(Gst���������8Xv��e0c{�`�)��Z�l-�>t�"uI����z������=��*!S�7�q��ck��N��\�ʰC�7��+ud�l=�P�k�-5|�lh�.;�����xXo�*�"���Rݾ�Q][n�Xԉ�^���)�	L���P������Cǁ]�Y�BMr0�t'}�rmw�H��W��w��%SqUHX�Rv��p�F|wQ��T����]s1FJ���a؞�F�lh��<ZWl�u�k�jE��y�s�@��}��g�M��\�! ����\���k7���ct��vk�o�y!߂��'k���j(z��btO�m��3�Yōة����:UB��e�֮�,J�#�KG,�<E�Jq,SP`t�h�������&�ث����XW�Ed�R	�{8�1hJ9}5���hѵ�=:w�n썌�ؖ�ktm/�{7�qrnތ|IH�4dۛ����8�T�
{w�4]�IQX5��]��bB��$\B��|~�:�rF��T��4��(�!	�����ņݒ,Ѵ���:��.[��q]�{5�[h�/KOv�Y�r.c?@��5��8�8Kjڏy\�V��?z�c�Ÿb��얕�mƱ;�#��V532w�6�K&.��\94��{�(��mW�#^�E���2�&dC{��q9?�QE�eM�ݱ�QW}�M74�$^��Fi/������S��t/(.�:Ӹ1	<�)�5X�*>r1�����z��*%���Xu� *Q�ɣ[�������������)5      g   8   x���� ���0��A������|l%����Z�.�lT�l�k�V�A����} "��      h   �   x���1
�@D뿧�+��f�{�$�J���Z��[��1B $g���{���0o�"<0�����
W6Z��0�Ҋ�V%�X/�
N���Oإ��Y�JW�/��mEJ�པ��29r�0��L+c�&Z%�"hibie�F��`�v!͓�Gӄ;Z��x.��m����e�2����sg      i   �   x�e��J�PF�O��0wn���-�����(!��Ԋ�6>�b���3L���[X}��;F��(
��=�{���l;����v�0=ا�C? !�y�9�������a9۠���5�vh�N.��(�5aH��Q��yO�>���vݔ}7�`o���h�9���9y�=G��$�`�����V7��Cv2��I�ω4�B����9�.�Q�      !    l           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            m           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            n           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                       false            o           1262    26138    g3    DATABASE     �   CREATE DATABASE g3 WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'Russian_Russia.1251' LC_CTYPE = 'Russian_Russia.1251';
    DROP DATABASE g3;
             postgres    false                        2615    26139    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
             postgres    false            p           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                  postgres    false    7            �            1259    26140    map    TABLE     B   CREATE TABLE public.map (
    id text NOT NULL,
    chunk text
);
    DROP TABLE public.map;
       public         postgres    false    7            �            1259    26146    masterships    TABLE     �   CREATE TABLE public.masterships (
    id integer NOT NULL,
    name text,
    short_name text,
    min integer,
    max integer
);
    DROP TABLE public.masterships;
       public         postgres    false    7            �            1259    26152    mastery_items    TABLE     +  CREATE TABLE public.mastery_items (
    id integer NOT NULL,
    mastery text,
    category text,
    ingredient text,
    name text,
    rarity integer,
    areas text,
    area_size integer,
    min integer,
    max integer,
    is_countable boolean,
    is_liquid boolean,
    "limit" integer
);
 !   DROP TABLE public.mastery_items;
       public         postgres    false    7            �            1259    26158    params    TABLE     =   CREATE TABLE public.params (
    key text,
    value text
);
    DROP TABLE public.params;
       public         postgres    false    7            �            1259    26164    person_inventory    TABLE     �   CREATE TABLE public.person_inventory (
    id text NOT NULL,
    person_id integer NOT NULL,
    item_id integer NOT NULL,
    weight text,
    quality text,
    creation_date integer,
    exp_date integer,
    is_deleted boolean
);
 $   DROP TABLE public.person_inventory;
       public         postgres    false    7            �            1259    26170    person_masterships    TABLE     |   CREATE TABLE public.person_masterships (
    person_id integer NOT NULL,
    mastery_id integer NOT NULL,
    skill text
);
 &   DROP TABLE public.person_masterships;
       public         postgres    false    7            �            1259    26176    persons    TABLE     �   CREATE TABLE public.persons (
    id integer NOT NULL,
    name text,
    age integer,
    is_male boolean,
    chunk_id text,
    day_action text
);
    DROP TABLE public.persons;
       public         postgres    false    7            �            1259    26214    t_user    TABLE     �   CREATE TABLE public.t_user (
    user_id bigint NOT NULL,
    username character varying,
    first_name character varying,
    last_name character varying,
    lang character varying,
    creation_date timestamp without time zone NOT NULL
);
    DROP TABLE public.t_user;
       public         postgres    false    7            b          0    26140    map 
   TABLE DATA               (   COPY public.map (id, chunk) FROM stdin;
    public       postgres    false    181   <       c          0    26146    masterships 
   TABLE DATA               E   COPY public.masterships (id, name, short_name, min, max) FROM stdin;
    public       postgres    false    182   R       d          0    26152    mastery_items 
   TABLE DATA               �   COPY public.mastery_items (id, mastery, category, ingredient, name, rarity, areas, area_size, min, max, is_countable, is_liquid, "limit") FROM stdin;
    public       postgres    false    183   �       e          0    26158    params 
   TABLE DATA               ,   COPY public.params (key, value) FROM stdin;
    public       postgres    false    184   4       f          0    26164    person_inventory 
   TABLE DATA               x   COPY public.person_inventory (id, person_id, item_id, weight, quality, creation_date, exp_date, is_deleted) FROM stdin;
    public       postgres    false    185   [       g          0    26170    person_masterships 
   TABLE DATA               J   COPY public.person_masterships (person_id, mastery_id, skill) FROM stdin;
    public       postgres    false    186   �       h          0    26176    persons 
   TABLE DATA               O   COPY public.persons (id, name, age, is_male, chunk_id, day_action) FROM stdin;
    public       postgres    false    187   4       i          0    26214    t_user 
   TABLE DATA               _   COPY public.t_user (user_id, username, first_name, last_name, lang, creation_date) FROM stdin;
    public       postgres    false    188   �       �           2606    26183    map map_id_key 
   CONSTRAINT     G   ALTER TABLE ONLY public.map
    ADD CONSTRAINT map_id_key UNIQUE (id);
 8   ALTER TABLE ONLY public.map DROP CONSTRAINT map_id_key;
       public         postgres    false    181            �           2606    26185    masterships masterships_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.masterships
    ADD CONSTRAINT masterships_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.masterships DROP CONSTRAINT masterships_pkey;
       public         postgres    false    182            �           2606    26187     mastery_items mastery_items_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.mastery_items
    ADD CONSTRAINT mastery_items_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.mastery_items DROP CONSTRAINT mastery_items_pkey;
       public         postgres    false    183            �           2606    26189    params params_key_key 
   CONSTRAINT     O   ALTER TABLE ONLY public.params
    ADD CONSTRAINT params_key_key UNIQUE (key);
 ?   ALTER TABLE ONLY public.params DROP CONSTRAINT params_key_key;
       public         postgres    false    184            �           2606    26191 (   person_inventory person_inventory_id_key 
   CONSTRAINT     a   ALTER TABLE ONLY public.person_inventory
    ADD CONSTRAINT person_inventory_id_key UNIQUE (id);
 R   ALTER TABLE ONLY public.person_inventory DROP CONSTRAINT person_inventory_id_key;
       public         postgres    false    185            �           2606    26193    persons persons_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.persons
    ADD CONSTRAINT persons_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.persons DROP CONSTRAINT persons_pkey;
       public         postgres    false    187            �           1259    26220    t_user_user_id_idx    INDEX     O   CREATE UNIQUE INDEX t_user_user_id_idx ON public.t_user USING btree (user_id);
 &   DROP INDEX public.t_user_user_id_idx;
       public         postgres    false    188            �           2606    26194 $   person_inventory person_inventory_fk    FK CONSTRAINT     �   ALTER TABLE ONLY public.person_inventory
    ADD CONSTRAINT person_inventory_fk FOREIGN KEY (person_id) REFERENCES public.persons(id);
 N   ALTER TABLE ONLY public.person_inventory DROP CONSTRAINT person_inventory_fk;
       public       postgres    false    2026    185    187            �           2606    26199 &   person_inventory person_inventory_fk_1    FK CONSTRAINT     �   ALTER TABLE ONLY public.person_inventory
    ADD CONSTRAINT person_inventory_fk_1 FOREIGN KEY (item_id) REFERENCES public.mastery_items(id);
 P   ALTER TABLE ONLY public.person_inventory DROP CONSTRAINT person_inventory_fk_1;
       public       postgres    false    185    2020    183            �           2606    26204 (   person_masterships person_masterships_fk    FK CONSTRAINT     �   ALTER TABLE ONLY public.person_masterships
    ADD CONSTRAINT person_masterships_fk FOREIGN KEY (person_id) REFERENCES public.persons(id);
 R   ALTER TABLE ONLY public.person_masterships DROP CONSTRAINT person_masterships_fk;
       public       postgres    false    186    2026    187            �           2606    26209 *   person_masterships person_masterships_fk_1    FK CONSTRAINT     �   ALTER TABLE ONLY public.person_masterships
    ADD CONSTRAINT person_masterships_fk_1 FOREIGN KEY (mastery_id) REFERENCES public.masterships(id);
 T   ALTER TABLE ONLY public.person_masterships DROP CONSTRAINT person_masterships_fk_1;
       public       postgres    false    186    182    2018           