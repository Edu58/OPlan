defmodule Oplan.Event.EventComment do
  use Ecto.Schema
  import Ecto.Changeset

  alias Oplan.Account.User
  alias Oplan.Event.Event

  schema "event_comments" do
    field :comment, :string
    belongs_to :user, User
    belongs_to :event, Event

    timestamps()
  end

  @doc false
  def changeset(event_comment, attrs) do
    event_comment
    |> cast(attrs, [:comment, :user_id, :event_id])
    |> validate_required([:comment, :user_id, :event_id])
  end
end
