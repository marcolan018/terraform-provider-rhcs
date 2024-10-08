variable "idp_url" {
  type = string
}
// Shared by all of the IDPs
variable "cluster_id" {
  type = string
}
variable "name" {
  type = string
}
variable "mapping_method" {
  type    = string
  default = "claim"
}
variable "bind_dn" {
  type    = string
  default = null
}

variable "bind_password" {
  type      = string
  default   = null
  sensitive = true
}

variable "insecure" {
  type    = bool
  default = false
}

variable "attributes" {
  type = object({
    email              = optional(list(string))
    id                 = optional(list(string))
    name               = optional(list(string))
    preferred_username = optional(list(string))
  })
  default     = null
  description = "attributes optional fields list"
}

variable "ca" {
  type    = string
  default = null
}