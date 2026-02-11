resource "waldur_openstack_volume_attachment" "example" {
  volume   = waldur_openstack_volume.example.id
  instance = waldur_openstack_instance.example.url
}
