import 'dart:io';
import 'package:flutter/material.dart';
import 'pages/login_page.dart';
import 'pages/add_friend.dart';
import 'package:image_picker/image_picker.dart';

class Todo {
  final int id;
  final String name;
  final String img;
  final String msg;

  const Todo(this.id, this.name, this.img, this.msg);
}

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primarySwatch: Colors.indigo,
      ),
      home: const MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});
  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  int _selectedIndex = 0;
  static final List<Widget> _widgetOptions = <Widget>[
    const MainPage(
      todos: [
        Todo(
            1,
            "Uno",
            "https://eloutput.com/wp-content/uploads/2019/05/uno-cartas.jpg",
            "Tremendo juegardo"
        ),
        Todo(
            2,
            "Ajedrez",
            "https://images.chesscomfiles.com/uploads/v1/blog/530042.33c6be72.5000x5000o.e96f2c4df196.jpeg",
            "Tremendo juegardo"),
        Todo(
            3,
            "Póker",
            "https://crehana-blog.imgix.net/media/filer_public/8c/a4/8ca49656-e762-45fc-81e0-948f1e7bc9c3/as-poker.jpeg?auto=format&q=50",
            "Tremendo juegardo"
        ),
        Todo(
            4,
            "GTA 5",
            "https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fgrand-theft-auto-v%2Fhome%2FGTAV_EGS_Artwork_1920x1080_Hero-Carousel_V06-1920x1080-1503e4b1320d5652dd4f57466c8bcb79424b3fc0.jpg",
            "Sí, incluso aquí lo puedes jugar"
        ),
      ],
    ),
    const FriendsPage(),
    const ProfilePage(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('App no tan mierdosa'),
        actions: [
          PopupMenuButton(
            itemBuilder: (context) {
              return [
                PopupMenuItem(
                  value: 0,
                  child: Text("Ajustes"),
                ),
                PopupMenuItem(
                  value: 1,
                  child: Text("Cerrar sesión"),
                ),
              ];
            },
          ),
        ],
      ),
      body: Center(
        child: _widgetOptions.elementAt(_selectedIndex),
      ),
      bottomNavigationBar: BottomNavigationBar(
        items: const [
          BottomNavigationBarItem(
            tooltip: 'Juegos',
            icon: Icon(Icons.games_rounded),
            label: 'Juegos',
          ),
          BottomNavigationBarItem(
            tooltip: 'Amigos',
            icon: Icon(Icons.people_alt),
            label: 'Amigos',
          ),
          BottomNavigationBarItem(
            tooltip: 'Perfil',
            icon: Icon(Icons.account_circle),
            label: 'Perfil',
          ),
        ],
        iconSize: 27,
        backgroundColor: Colors.indigo,
        showUnselectedLabels: false,
        selectedItemColor: Colors.amber,
        unselectedItemColor: Colors.white,
        onTap: _onItemTapped,
        currentIndex: _selectedIndex,
      ),
      drawer: Drawer(
        child: Container(
          color: Colors.indigo,
          padding: const EdgeInsets.symmetric(vertical: 0, horizontal: 10),
          child: ListView.separated(
            itemCount: 50,
            itemBuilder: (_, index) {
              return DefaultTextStyle(
                style: const TextStyle(color: Colors.white),
                child: SizedBox(
                  height: 50,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceAround,
                    children: const [
                      Icon(Icons.home, color: Colors.white),
                      Text('Esto es una prueba lololololololo')
                    ],
                  ),
                ),
              );
            },
            separatorBuilder: (context, index) => Container(
              height: 1,
              color: Colors.white,
            ),
          ),
        ),
      ),
    );
  }
}

class MainPage extends StatefulWidget {
  const MainPage({super.key, required this.todos});
  final List<Todo> todos;
  @override
  State<MainPage> createState() => _MainPageState();
}

class _MainPageState extends State<MainPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListView.builder(
        itemCount: widget.todos.length,
        itemBuilder: (context, index) {
          return IconButton(
            icon: Stack(
              alignment: Alignment.bottomLeft,
              children: [
                Hero(
                  tag: widget.todos[index].id,
                  child: Image.network(
                    widget.todos[index].img,
                    width: 1000,
                    height: 500,
                    fit: BoxFit.cover,
                  ),
                ),
                Container(
                  decoration: BoxDecoration(
                    color: Colors.white,
                    gradient: LinearGradient(
                      begin: FractionalOffset.topCenter,
                      end: FractionalOffset.bottomCenter,
                      colors: [
                        Colors.black.withOpacity(0.0),
                        Colors.black,
                      ],
                    ),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(10),
                  child: Text(
                    widget.todos[index].name,
                    style: const TextStyle(
                      fontSize: 20,
                      color: Colors.white,
                    ),
                  ),
                ),
              ],
            ),
            iconSize: 250,
            onPressed: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => GamePage(todo: widget.todos[index]),
                ),
              );
            },
          );
        },
      ),
    );
  }
}

class GamePage extends StatelessWidget {
  const GamePage({super.key, required this.todo});
  final Todo todo;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(todo.name),
      ),
      body: SingleChildScrollView(
        child: Column(
          children: [
            Hero(
              tag: todo.id,
              child: Image.network(todo.img, fit: BoxFit.cover),
            ),
            Container(
              color: Colors.indigo,
              child: Align(
                child: Text(
                  todo.msg,
                  style: const TextStyle(color: Colors.white, fontSize: 40),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

class FriendsPage extends StatefulWidget {
  const FriendsPage({super.key});
  @override
  State<FriendsPage> createState() => _FriendsPageState();
}

class _FriendsPageState extends State<FriendsPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListView.separated(
        itemCount: 20,
        itemBuilder: (context, index) {
          return ListTile(
            leading: FriendListTile(),
            title: Row(
              children: const [
                Text(
                  'Amigo falso',
                  style: TextStyle(
                    color: Colors.blueGrey,
                    fontSize: 18,
                  ),
                ),
              ],
            ),
            subtitle: Text('perro'),
            onTap: () {},
          );
        },
        separatorBuilder: (context, index) => Container(
          height: 1,
          color: Colors.indigoAccent,
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          openDialog();
        },
        tooltip: 'Añadir amigo',
        child: const Icon(Icons.add_rounded, size: 30),
      ),
    );
  }

  Future openDialog() => showDialog(
    context: context,
    builder: (context) => Dialog(
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(10),
      ),
      elevation: 0,
      backgroundColor: Colors.transparent,
      child: AddFriendPage(),
    ),
  );
}

class FriendListTile extends StatelessWidget {
  const FriendListTile({super.key});

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 100,
      child: Stack(
        children: [
          CircleAvatar(
            backgroundImage: ResizeImage(
              AssetImage('images/pepoclown.jpg'),
              width: 120, height: 120
            ),
            radius: 30,
          ),
          Container(
            height: 60,
            width: 60,
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              border: Border.all(
                color: Colors.indigoAccent,
                width: 3.0,
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class ProfilePage extends StatefulWidget {
  const ProfilePage({super.key});
  @override
  State<ProfilePage> createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: Column(
          children: [
            const Profile(),
            Container(
              padding: const EdgeInsets.all(20),
              color: Colors.indigo[50],
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  Column(
                    children: [
                      Text(
                        'PARTIDAS JUGADAS',
                        style: TextStyle(
                            color: Colors.blueGrey[700],
                            fontWeight: FontWeight.bold),
                      ),
                      Text(
                        '420',
                        style: TextStyle(
                          color: Colors.blueGrey,
                          fontSize: 20,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                    ],
                  ),
                  Container(
                    width: 1,
                    height: 50,
                    color: Colors.indigoAccent,
                  ),
                  Column(
                    children: [
                      Text(
                        'WINRATE',
                        style: TextStyle(
                          color: Colors.blueGrey[700],
                          fontWeight: FontWeight.bold
                        ),
                      ),
                      Text(
                        '100%',
                        style: TextStyle(
                          color: Colors.blueGrey,
                          fontSize: 20,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                    ],
                  ),
                  Container(
                    width: 1,
                    height: 50,
                    color: Colors.indigoAccent,
                  ),
                  Column(
                    children: [
                      Text(
                        'SUERTE',
                        style: TextStyle(
                          color: Colors.blueGrey[700],
                          fontWeight: FontWeight.bold
                        ),
                      ),
                      Text(
                        '9000',
                        style: TextStyle(
                          color: Colors.blueGrey,
                          fontSize: 20,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                    ],
                  ),
                ],
              ),
            ),
            ListView.separated(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
              itemCount: 20,
              itemBuilder: (_, index) {
                return DefaultTextStyle(
                  style: const TextStyle(color: Colors.white),
                  child: SizedBox(
                    child: Row(
                      children: const [
                        Padding(
                          padding: EdgeInsets.all(10),
                          child: Icon(
                            Icons.star,
                            color: Colors.indigoAccent,
                            size: 35,
                          ),
                        ),
                        Text(
                          'Victoria magistral',
                          style: TextStyle(
                            color: Colors.black,
                          ),
                        ),
                        SizedBox(
                          width: 10,
                        ),
                        Icon(Icons.thumb_up),
                      ],
                    ),
                  ),
                );
              },
              separatorBuilder: (context, index) =>
                  Container(height: 1, color: Colors.indigoAccent),
            ),
          ],
        ),
      ),
    );
  }
}

class Profile extends StatefulWidget {
  const Profile({super.key});

  @override
  State<Profile> createState() => _ProfileState();
}

class _ProfileState extends State<Profile> {
  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Container(
          margin: const EdgeInsets.all(10),
          height: 100,
          width: 100,
          decoration: BoxDecoration(
            shape: BoxShape.circle,
            border: Border.all(
              color: Colors.indigoAccent,
              width: 3.0,
            ),
          ),
          child: IconButton(
            padding: const EdgeInsets.all(0),
            onPressed: () {
              Navigator.push(
                context,
                MaterialPageRoute(builder: (context) => const ProfilePicture()),
              );
            },
            icon: const Hero(
              tag: 'foto',
              child: CircleAvatar(
                backgroundImage: ResizeImage(
                    AssetImage('images/pepoclown.jpg'),
                    width: 200,
                    height: 200,
                ),
                radius: 50,
              ),
            ),
          ),
        ),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: const [
            Text(
              'Ismaber#1234',
              style: TextStyle(
                color: Colors.indigoAccent,
                fontSize: 25,
                fontWeight: FontWeight.bold,
              ),
            ),
            SizedBox(
              height: 10,
            ),
            Text(
              'El puto amo',
              style: TextStyle(
                color: Colors.blueGrey,
                fontSize: 18,
                fontWeight: FontWeight.w300,
              ),
            ),
          ],
        ),
      ],
    );
  }
}

class ProfilePicture extends StatefulWidget {
  const ProfilePicture({super.key});

  @override
  State<ProfilePicture> createState() => _ProfilePictureState();
}

class _ProfilePictureState extends State<ProfilePicture> {
  File? image;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Perfil'),
      ),
      body: SingleChildScrollView(
        child: Column(
          children: [
            Container(
              height: 300,
              width: double.infinity,
              color: Colors.indigo[50],
              child: image == null
                ? IconButton(
                  padding: const EdgeInsets.all(0),
                  onPressed: () {
                    openDialog();
                  },
                  iconSize: 300,
                  icon: Hero(
                    tag: 'foto',
                    child: CircleAvatar(
                      backgroundImage: ResizeImage(
                        AssetImage('images/pepoclown.jpg'),
                        width: 300,
                        height: 300,
                      ),
                      radius: 100,
                    ),
                  ),
                )
                : IconButton(
                  padding: const EdgeInsets.all(0),
                  onPressed: () {
                    openDialog();
                  },
                  iconSize: 300,
                  icon: Hero(
                    tag: 'foto',
                    child: CircleAvatar(
                      backgroundImage: FileImage(image!),
                      radius: 100,
                    ),
                  ),
                ),
            ),
            Container(
              width: double.infinity,
              padding: const EdgeInsets.all(10),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Text(
                    'Nickname',
                    style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  TextField(
                    keyboardType: TextInputType.text,
                    decoration: InputDecoration(
                      border: InputBorder.none,
                      hintText: 'Ismaber',
                    ),
                  ),
                  Container(
                    color: Colors.indigoAccent,
                    height: 1,
                  ),
                  const Text(
                    'Descripción',
                    style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  TextField(
                    keyboardType: TextInputType.text,
                    decoration: InputDecoration(
                      border: InputBorder.none,
                      hintText: 'Sobre mí',
                    ),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }

  Future openDialog() => showDialog(
        context: context,
        builder: (context) => Dialog(
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(10),
          ),
          elevation: 0,
          backgroundColor: Colors.transparent,
          child: SingleChildScrollView(
            child: Container(
              padding: const EdgeInsets.all(15),
              decoration: BoxDecoration(
                  color: Colors.white, borderRadius: BorderRadius.circular(10)),
              child: Column(
                children: [
                  const Text(
                    'Elegir foto de perfil',
                    style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  SizedBox(
                    height: 50,
                    width: double.infinity,
                    child: ElevatedButton(
                      onPressed: () {
                        _pickImageFromGallery();
                      },
                      child: const Text('Galería'),
                    ),
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  SizedBox(
                    height: 50,
                    width: double.infinity,
                    child: ElevatedButton(
                      onPressed: () {
                        _pickImageFromCamera();
                      },
                      child: const Text('Cámara'),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      );

  _pickImageFromGallery() async {
    XFile? pickedFile = await ImagePicker().pickImage(
      source: ImageSource.gallery,
      maxWidth: 1800,
      maxHeight: 1800,
    );
    if (pickedFile != null) {
      setState(() {
        image = File(pickedFile.path);
      });
    }
    if (context.mounted) Navigator.pop(context);
  }

  /// Get from Camera
  _pickImageFromCamera() async {
    XFile? pickedFile = await ImagePicker().pickImage(
      source: ImageSource.camera,
      maxWidth: 1800,
      maxHeight: 1800,
    );
    if (pickedFile != null) {
      setState(() {
        image = File(pickedFile.path);
      });
    }
    if (context.mounted) Navigator.pop(context);
  }
}