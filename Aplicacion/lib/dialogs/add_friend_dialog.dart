import 'package:flutter/material.dart';

final addFriendFormKey = GlobalKey<FormState>();

class AddFriendPage extends StatelessWidget {
  const AddFriendPage({super.key});

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Añadir amigo con ID'),
      content: Form(
        key: addFriendFormKey,
        child: TextFormField(
          keyboardType: TextInputType.text,
          decoration: const InputDecoration(
            border: OutlineInputBorder(),
            hintText: '#1234',
          ),
          validator: (value) {
            if (value == null || value.isEmpty) {
              return 'El campo es obligatorio';
            }
            return null;
          },
        ),
      ),
      actionsAlignment: MainAxisAlignment.spaceEvenly,
      actions: [
        FilledButton(
          onPressed: () {
            if(addFriendFormKey.currentState!.validate()) {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text('Amigo añadido'),
                  showCloseIcon: true,
                  closeIconColor: Colors.white,
                ),
              );
              Navigator.pop(context);
            }
          },
          child: const Text('Añadir amigo'),
        ),
        FilledButton(
          onPressed: () {
            Navigator.pop(context);
          },
          child: const Text('Cancelar'),
        ),
      ],
    );
  }
}