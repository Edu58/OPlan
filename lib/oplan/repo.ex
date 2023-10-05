defmodule Oplan.Repo do
  use Ecto.Repo,
    otp_app: :oplan,
    adapter: Ecto.Adapters.Postgres
end
