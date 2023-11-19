defmodule Oplan.Profiles.Profile do
  use Ecto.Schema
  import Ecto.Changeset

  alias Oplan.Account.User

  schema "profiles" do
    field :avatar, :string
    field :city, :string
    field :date_of_birth, :string
    field :nationality, :string
    field :phone_number, :string
    belongs_to :user, User

    timestamps()
  end

  @doc false
  def changeset(profile, attrs) do
    profile
    |> cast(attrs, [:avatar, :date_of_birth, :phone_number, :nationality, :city, :user])
  end
end
