// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { ACTIVATION_MODES, LORAWAN_VERSIONS, LORAWAN_PHY_VERSIONS } from '@console/lib/device-utils'

// End device selectors.

export const getActivationMode = device =>
  device.supports_join === true
    ? ACTIVATION_MODES.OTAA
    : device.multicast === true
    ? ACTIVATION_MODES.MULTICAST
    : device.supports_join === false && device.multicast === false
    ? ACTIVATION_MODES.ABP
    : ACTIVATION_MODES.NONE

export const getLorawanVersion = device => device.lorawan_version || '1.1.0'

export const getApplicationServerAddress = device => device.application_server_address
export const getNetworkServerAddress = device => device.network_server_address
export const getJoinServerAddress = device => device.join_server_address

// End device repository utils.

export const SELECT_OTHER_OPTION = '_other_'
export const isOtherOption = option => option === SELECT_OTHER_OPTION
export const hasSelectedDeviceRepositoryOther = version =>
  version && Object.values(version).some(value => isOtherOption(value))
export const hasCompletedDeviceRepositorySelection = version =>
  version && Object.values(version).every(value => value)
export const hasValidDeviceRepositoryType = (version, template) =>
  !hasSelectedDeviceRepositoryOther(version) &&
  hasCompletedDeviceRepositorySelection(version) &&
  Boolean(template)
export const mayProvisionDevice = (values, template) => {
  const { frequency_plan_id, lorawan_version, lorawan_phy_version, _inputMethod, version_ids } =
    values

  return (
    (_inputMethod === 'manual' &&
      Boolean(frequency_plan_id) &&
      Boolean(lorawan_version) &&
      Boolean(lorawan_phy_version)) ||
    (_inputMethod === 'device-repository' &&
      hasValidDeviceRepositoryType(version_ids, template) &&
      Boolean(frequency_plan_id))
  )
}

/*
  `hardware_version` is not required when registering an end device in the device repository, so for 
  certain end device models it can be missing. When this is the case, we still want to allow the users
  to select such models because `firmware_version` (that might depend on hw version) and `band_id`
  are required. `SELECT_UNKNOWN_HW_OPTION` option represents such end devices.
*/
export const SELECT_UNKNOWN_HW_OPTION = '_unknown_hw_version_'
export const isUnknownHwVersion = option => option === SELECT_UNKNOWN_HW_OPTION

export const REGISTRATION_TYPES = Object.freeze({
  SINGLE: 'single',
  MULTIPLE: 'multiple',
})

export const getLorawanVersionLabel = ({ lorawan_version }) => {
  const { label } = LORAWAN_VERSIONS.find(version => version.value === lorawan_version) || {}

  return label
}

export const getLorawanPhyVersionLabel = ({ lorawan_phy_version }) => {
  const { label } =
    LORAWAN_PHY_VERSIONS.find(version => version.value === lorawan_phy_version) || {}

  return label
}

export const DEVICE_CLASS_MAP = {
  CLASS_A: 'class-a',
  CLASS_B: 'class-b',
  CLASS_C: 'class-c',
  CLASS_B_C: 'class-b-c',
}
