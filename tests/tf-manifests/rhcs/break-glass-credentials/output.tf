output kubeconfig {
  value = rhcs_break_glass_credential.rosa_break_glass_credential.kubeconfig
  sensitive = true
}

output username {
  value = rhcs_break_glass_credential.rosa_break_glass_credential.username
}

output expiration_timestamp {
  value = rhcs_break_glass_credential.rosa_break_glass_credential.expiration_timestamp
}

output revocation_timestamp {
  value = rhcs_break_glass_credential.rosa_break_glass_credential.revocation_timestamp
}

output status {
  value = rhcs_break_glass_credential.rosa_break_glass_credential.status
}

