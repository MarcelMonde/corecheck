resource "google_storage_bucket" "bitcoin-coverage-cache" {
  name          = "bitcoin-coverage-cache"
  location      = "europe-west1"
  force_destroy = true
  versioning {
    enabled = false
  }
  public_access_prevention = "enforced"
}

resource "google_storage_bucket" "bitcoin-coverage-ccache" {
  name          = "bitcoin-coverage-ccache"
  location      = "europe-west1"
  force_destroy = true
  versioning {
    enabled = false
  }
  public_access_prevention = "enforced"
}

resource "google_storage_bucket" "bitcoin-coverage-data" {
  name          = "bitcoin-coverage-data"
  location      = "europe-west1"
  force_destroy = true
  versioning {
    enabled = false
  }
}

resource "google_storage_bucket_iam_member" "bitcoin-coverage-data-public" {
  bucket = google_storage_bucket.bitcoin-coverage-data.name
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}