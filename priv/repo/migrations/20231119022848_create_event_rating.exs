defmodule Oplan.Repo.Migrations.CreateEventRating do
  use Ecto.Migration

  def change do
    create table(:event_rating) do
      add :rating, :integer
      add :user_id, references(:users, on_delete: :nothing)
      add :event_id, references(:events, on_delete: :nothing)

      timestamps()
    end

    create unique_index(:event_rating, [:user_id, :event_id])
  end
end
