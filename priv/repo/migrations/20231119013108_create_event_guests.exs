defmodule Oplan.Repo.Migrations.CreateEventGuests do
  use Ecto.Migration

  def change do
    create table(:event_guests) do
      add :user_id, references(:users, on_delete: :nothing)
      add :event_id, references(:events, on_delete: :nothing)
      add :confirmed, :boolean, default: false
      add :arrived, :boolean, default: false
      add :departed, :boolean, default: false

      timestamps()
    end

    create unique_index(:event_guests, [:user_id, :event_id])
  end
end
