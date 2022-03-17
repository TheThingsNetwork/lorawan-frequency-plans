# LoRaWAN Frequency Plans for The Things Stack

This repository contains default frequency plans to be used with The Things Stack.

Frequency plans for The Things Stack contain channels, data rates and radio configuration, as well as settings to comply with regional regulations, i.e. time-off-air, dwell time and listen-before-talk.

Frequency plans are defined for a band. Bands are specified by the LoRa Alliance Technical Committee and are published as the [LoRaWAN Regional Parameters technical document](https://lora-alliance.org/resource-hub).

## File Format

Frequency plan are defined in YAML files. Most settings in the frequency plan are optional. When not specifying optional settings, the band defaults are used.

```yml
band-id: BAND_ID               # ID of the band (needs to match band-id in the index)
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
  band-id: EU_863_870      # ID of the LoRaWAN band (needs to match band-id in the definition)
  base-id: EU_863_870      # ID that this frequency plan extends (refers to id of another frequency plan)
  name: Region 863-870 MHz # Name of the frequency plan, ending with frequency ranges
  base-frequency: 868      # Base frequency in MHz for hardware support (433, 470, 868 or 915)
  country-codes: []        # List of 2-digit ISO country codes for countries where this plan can be used
  file: EU_863_870.yml     # File of the frqeuency plan definition
```

> Country codes are taken from the [LoRaWAN Regional Parameters 1.0.1 Specification](https://lora-alliance.org/sites/default/files/2020-06/rp_2-1.0.1.pdf)

## Contributing

Thank you for your interest in building this thing together with us. We're really happy with our active community and are glad that you're a part of it.

The Things Stack uses the `github.com/TheThingsNetwork/lorawan-frequency-plans` as default source for fetching frequency plans. Therefore, contributing to this open source repository makes frequency plans automatically available to Stack deployments with default settings. You can contribute by submitting pull requests. Are you new to GitHub? That's great! [Read here about pull requests](https://help.github.com/articles/about-pull-requests/). Please also use the editor settings as defined in `.editorconfig`.

### Local Regulations

When submitting a new frequency plan or making changes to an existing frequency plan, please make sure that the band is allowed to be used in the concerning region and that settings respect regional regulations. When submitting a pull request for a new region, please upload or link to a document that describes the local regulations.

### Frequency Plan Naming

The IDs of frequency plans in this repository follow the naming of the corresponding bands in the LoRaWAN Regional Parameters specification. For compatibility reasons, frequency plan IDs can not be modified at a later stage.

The name of the frequency plan should mention the region, band and important specifics about the frequency plan. The description can further explain where and how the frequency plan is used.

### Band ID

The Things Stack supports the following bands:

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

See the [LoRaWAN Regional Parameters](https://lora-alliance.org/resource-hub) for more details on the band definitions. If you want to contribute to the band definitions in The Things Stack, please contribute to the [TheThingsNetwork/lorawan-stack](https://github.com/TheThingsNetwork/lorawan-stack) repository.

### Spectrum Limitations

Regulatory restrictions for the band can be configured with the `max-eirp`, `duty-cycle`, `dwell-time`, `time-off-air` and `listen-before-talk` settings.

These restrictions can also be specified for sub-bands, which are defined by setting a minimum and maximum frequency, and by configuring the restrictions that apply between these frequencies.

### Radios and Channels

Gateways need to be configured with center frequencies for each radio. Typical concentrator boards have two radios, identified by index `0` and `1`. Each of these radios can handle four 125kHz uplink channels. For an optimal signal quality, these channels should be placed max 400kHz away from the center frequency of the radio. Each of these channels has a `frequency`, a `min-data-rate` and `max-data-rate`.

The `downlink-channels` define the downlink channels corresponding to channels defined under `uplink-channels`.

With `lora-standard-channel` a single fixed-SF uplink channel can be configured. This is typically a 250kHz channel. With `fsk-channel` a single GFSK channel can be configured.

### LoRaWAN Settings

Other LoRaWAN settings that can be configured in a frequency plan:

- RX2 parameters: `rx2-channel` and `rx2-default-data-rate`
- Class B ping slot: `ping-slot` and `ping-slot-default-data-rate`
