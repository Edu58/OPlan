defmodule Oplan.Event.EventType do
  use Ecto.Schema
  import Ecto.Changeset

  schema "event_types" do
    field :name, :string

    timestamps()
  end

  @doc false
  def changeset(event_type, attrs) do
    event_type
    |> cast(attrs, [:name])
    |> validate_required([:name])
  end
end
