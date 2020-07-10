# LoRaWAN Frequency Plans for The Things Stack

This repository contains default frequency plans to be used with The Things Network Stack for LoRaWAN.

Frequency plans for The Things Network Stack contain channels, data rates and radio configuration, as well as settings to comply with regional regulations, i.e. time-off-air, dwell time and listen-before-talk.

Frequency plans are defined for a band. Bands are specified by the LoRa Alliance Technical Committee and are published as the [LoRaWAN Regional Parameters technical document](https://lora-alliance.org/resource-hub).

## File Format

Frequency plan are defined in YAML files. Most settings in the frequency plan are optional. When not specifying optional settings, the band defaults are used.

```yml
band-id: BAND_ID               # ID of the band
sub-bands:
- min-frequency: 868000000     # Minimum frequency (Hz, inclusive)
  max-frequency: 868600000     # Maximum frequency (Hz, inclusive)
  duty-cycle: 0.01             # Duty cycle for this sub-band (optional; default: 1)
  max-eirp: 16.15              # Maximum EIRP for this sub-band (optional; takes precedence over frequency plan's max-eirp)
uplink-channels:               # List of uplink channels (zero indexed)
- frequency: 868100000         # Frequency (Hz)
  min-data-rate: 0             # Mininum data rate index
  max-data-rate: 5             # Maximum data rate index
  radio: 0                     # Radio index (see below)
downlink-channels:             # List of downlink channels (zero indexed)
- frequency: 868100000
  min-data-rate: 0
  max-data-rate: 5
  radio: 0
lora-standard-channel:         # LoRa standard channel (optional)
  frequency: 863000000
  data-rate: 6
  radio: 0
fsk-channel:                   # FSK channel (optional)
  frequency: 868800000
  data-rate: 7
  radio: 0
time-off-air:                  # Time-off-air (optional)
  fraction: 0.1                # Minimum fraction of the emission time (optional)
  duration: 1s                 # Minimum duration (optional)
dwell-time:                    # Dwell time (optional)
  uplinks: true                # Enabled for uplink (optional)
  downlinks: true              # Enabled for downlink (optional)
  duration: 1s                 # Duration (optional)
listen-before-talk:            # Listen-before-talk (optional)
  rssi-offset: 0               # RSSI offset (dbm)
  rssi-target: -80             # RSSI target (dbm)
  scan-time: 128000            # Scan time (nanoseconds)
radios:                        # Radio configuration (zero indexed, optional)
- enable: true                 # Enable the radio
  chip-type: SX1257            # Chip type
  frequency: 867500000         # Frequency (Hz)
  rssi-offset: -166            # RSSI offset (dbm)
  tx:                          # Radio transmission configuration (optional)
    min-frequency: 863000000   # Minimum frequency (Hz)
    max-frequency: 867000000   # Maximum frequency (Hz)
    notch-frequency: 129000    # Notch frequency 126000..250000 (Hz)
clock-source: 0                # Gateway clock source
ping-slot:                     # Class B ping slot settings (optional)
  frequency: 869525000
  min-data-rate: 0
  max-data-rate: 5
  radio: 0
ping-slot-default-data-rate: 3 # Default data rate index of class B ping slot (optional)
rx2-channel:                   # Rx2 channel (optional)
  frequency: 869525000
  min-data-rate: 0
  max-data-rate: 5
  radio: 0
rx2-default-data-rate: 0       # Default data rate index of Rx2 (optional)
max-eirp: 29.15                # Maximum EIRP (optional; used when sub-bands do not have max-eirp, takes precedence over band's default)
```

An index of frequency plans is in `frequency-plans.yml`:

```yml
- id: EU_863_870_TTN       # ID of the frequency plan
  base-id: EU_863_870      # ID that this frequency plan extends
  name: Region 863-870 MHz # Name of the frequency plan, ending with frequency ranges
  base-frequency: 868      # Base frequency in MHz for hardware support (433, 470, 868 or 915)
  country-codes: []        # List of 2-digit ISO country codes for countries where this plan can be used
  file: EU_863_870.yml     # File of the frqeuency plan definition
```

## Contributing

Thank you for your interest in building this thing together with us. We're really happy with our active community and are glad that you're a part of it.

The Things Network Stack uses the `github.com/TheThingsNetwork/lorawan-frequency-plans` as default source for fetching frequency plans. Therefore, contributing to this open source repository makes frequency plans automatically available to Stack deployments with default settings. You can contribute by submitting pull requests. Are you new to GitHub? That's great! [Read here about pull requests](https://help.github.com/articles/about-pull-requests/). Please also use the editor settings as defined in `.editorconfig`.

When submitting a new frequency plan or making changes to an existing frequency plan, please make sure that the band is allowed to be used in the concerning region and that settings respect regional regulations.

The Things Network Stack supports at least the following bands:

- `AS_923`: Asia 923 MHz
- `AU_915_928`: Australia 915 - 928 MHz
- `CN_470_510`: China 470 - 510 MHz
- `CN_779_787`: China 779 - 787 MHz
- `EU_433`: Europe 433 MHz
- `EU_863_870`: Europe 863 - 870 MHz
- `IN_865_870`: India 865 - 867 MHz
- `KR_920_923`: Korea 920 - 923 MHz
- `RU_864_870`: Russia 864 - 870 MHz
- `US_902_928`: United States 902 - 928 MHz

See the [LoRaWAN Regional Parameters](https://lora-alliance.org/resource-hub) for more band definitions. If you want to contribute to the band settings, please contribute to [The Things Network Stack for LoRaWAN](https://github.com/TheThingsNetwork/lorawan-stack).
