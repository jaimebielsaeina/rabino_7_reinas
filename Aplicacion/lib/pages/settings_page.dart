import 'package:flutter/material.dart';

class SettingsPage extends StatefulWidget {
  const SettingsPage({super.key});

  @override
  State<SettingsPage> createState() => _SettingsPageState();
}

class _SettingsPageState extends State<SettingsPage> {
  final Icon _muted = const Icon(
    Icons.volume_off,
  );
  final Icon _nonMuted = const Icon(
    Icons.volume_up,
  );

  double _musicValue = 100;
  double _oldMusicValue = 100;
  bool _muteMusic = false;
  late Icon _musicIcon = _nonMuted;

  double _soundEffectsValue = 100;
  double _oldSoundEffectsValue = 100;
  bool _muteSoundEffects = false;
  late Icon _soundEffectsIcon = _nonMuted;

  void muteUnmuteMusic() {
    setState(() {
      _muteMusic = !_muteMusic;
      if (_muteMusic) {
        _oldMusicValue = _musicValue;
        _musicValue = 0;
        _musicIcon = _muted;
      } else {
        print(_oldMusicValue);
        _musicValue = _oldMusicValue;
        _musicIcon = _nonMuted;
      }
    });
  }

  void muteUnmuteSoundEffects() {
    setState(() {
      _muteSoundEffects = !_muteSoundEffects;
      if (_muteSoundEffects) {
        _oldSoundEffectsValue = _soundEffectsValue;
        _soundEffectsValue = 0;
        _soundEffectsIcon = _muted;
      } else {
        _soundEffectsValue = _oldSoundEffectsValue;
        _soundEffectsIcon = _nonMuted;
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Ajustes'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(10),
        child: Column(
          children: [
            Row(
              children: [
                Container(
                  width: 80,
                  child: Text('Música'),
                ),
                Expanded(
                  child: Slider.adaptive(
                    value: _musicValue,
                    max: 100,
                    divisions: 100,
                    label: _musicValue.round().toString(),
                    activeColor: Colors.indigoAccent,
                    onChanged: (double value) {
                      if (_muteMusic) {
                        null;
                      } else {
                        setState(() {
                          _musicValue = value;
                        });
                      }
                    },
                  ),
                ),
                IconButton(
                  color: Colors.indigoAccent,
                  iconSize: 30,
                  icon: _musicIcon,
                  onPressed: () {muteUnmuteMusic();},
                ),
              ],
            ),
            Row(
              children: [
                Container(
                  width: 80,
                  child: Text('Efectos de sonido'),
                ),
                Expanded(
                  child: Slider.adaptive(
                    value: _soundEffectsValue,
                    max: 100,
                    divisions: 100,
                    label: _soundEffectsValue.round().toString(),
                    activeColor: Colors.indigoAccent,
                    onChanged: (double value) {
                      if (_muteSoundEffects) {
                        null;
                      } else {
                        setState(() {
                          _soundEffectsValue = value;
                        });
                      }
                    },
                  ),
                ),
                IconButton(
                  color: Colors.indigoAccent,
                  iconSize: 30,
                  icon: _soundEffectsIcon,
                  onPressed: (){muteUnmuteSoundEffects();},
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}